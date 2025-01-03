package global

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Logger *zap.SugaredLogger
	Client = resty.New()
)
