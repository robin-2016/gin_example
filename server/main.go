package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/robin-2016/gin_example/server/configs"
	"github.com/robin-2016/gin_example/server/global"
	"github.com/robin-2016/gin_example/server/pkg/logs"
	"github.com/robin-2016/gin_example/server/router"
)

func main() {
	configs.InitConfig()           //初始化配置文件
	global.Logger = logs.InitZap() //初始化日志
	global.DB = configs.InitDB()   //初始化数据库连接

	//初始化表
	if global.DB != nil {
		configs.MigrateDB()
		db, _ := global.DB.DB()
		// 关闭数据库连接
		defer db.Close()
	}

	r := router.InitRouter() //初始化路由

	// r.Run(configs.AppConfig.AppPort)
	srv := &http.Server{
		Addr:    configs.AppConfig.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	global.Logger.Infof("server listen on %v", configs.AppConfig.Port)
	<-quit
	global.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Errorf("Server Shutdown:", err)
	}
	global.Logger.Info("Server exiting")
}
