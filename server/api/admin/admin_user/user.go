package adminuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robin-2016/gin_example/server/global"
	"github.com/robin-2016/gin_example/server/model"
)

func AddUser(c *gin.Context) {
	var user model.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := global.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	global.Logger.Infof("user %v add succ", user.UserName)
}
