package tests

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

func Test_Logger(t *testing.T) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	newLogger.Info(ctx, ">>>>>>>>>>>>>>>日志测试>>>>>>>>>>>>>>>>>>>>>")
}
