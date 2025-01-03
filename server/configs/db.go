package configs

import (
	"os"

	"github.com/robin-2016/gin_example/server/global"
	"github.com/robin-2016/gin_example/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=" + AppConfig.DB.Host + " user=" + AppConfig.DB.User + " password=" + AppConfig.DB.Password + " dbname=" + AppConfig.DB.Name + " port=" + AppConfig.DB.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		global.Logger.Errorf("Failed to init db connection,error: %v", err)
	}
	return db
}

func MigrateDB() {
	err := global.DB.AutoMigrate(
		&model.Users{},
	)
	if err != nil {
		global.Logger.Errorf("Failed to migrate db,error: %v", err)
		os.Exit(0)
	}
}
