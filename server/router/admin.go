package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	adminuser "github.com/robin-2016/gin_example/server/api/admin/admin_user"
	"github.com/robin-2016/gin_example/server/configs"
	"github.com/robin-2016/gin_example/server/global"
)

func adminRLoad(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			global.Logger.Infof("test %v", "aaa")
			c.JSON(http.StatusOK, gin.H{
				"msg":  "pong",
				"port": configs.AppConfig.Port,
			})
		})
		api.POST("/user/add", adminuser.AddUser)
	}
}
