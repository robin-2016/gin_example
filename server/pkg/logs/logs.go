package logs

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/robin-2016/gin_example/server/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitZap() *zap.SugaredLogger {
	level, err := zapcore.ParseLevel(configs.AppConfig.Log.Level)
	if err != nil {
		log.Fatalf("Log level is err: %v", err)
	}
	core := zapcore.NewCore(zapEncoder(), zapcore.NewMultiWriteSyncer(zapWriter(), zapcore.AddSync(os.Stdout)), level)
	logger := zap.New(core).Sugar()
	defer logger.Sync()
	return logger
}

func zapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriter() zapcore.WriteSyncer {
	currentDir, _ := os.Getwd()
	strSeparator := string(filepath.Separator)
	logFile := currentDir + strSeparator + "logs" + strSeparator + time.Now().Format(time.DateOnly) + ".log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    configs.AppConfig.Log.MaxSize,
		MaxBackups: configs.AppConfig.Log.MaxBackups,
		MaxAge:     configs.AppConfig.Log.MaxAge,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
