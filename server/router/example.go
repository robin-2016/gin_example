package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robin-2016/gin_example/server/global"
)

func miniRLoad(r *gin.Engine) {
	userGroup := r.Group("/example")
	{
		userGroup.GET("/info", func(c *gin.Context) {
			global.Logger.Info("example router")
			c.JSON(http.StatusOK, gin.H{"msg": "example"})
		})
	}
}
