package config

import (
	"github.com/jinzhu/configor"
)

func Load(configLocation string) {
	// 加载配置文件
	configor.Load(&Config, configLocation)
}

var Config = struct {
	Server struct {
		Appname string `default:"jy"`
		Port    string `default:":5000"`
	}

	DB struct {
		Enable   bool   `default:"false"`
		Driver   string `default:"mysql"`
		Host     string `deault:"127.0.0.1"`
		Dbname   string `default:"zld"`
		Port     string `deault:"3306"`
		Username string `default:"root"`
		Password string `default:"password"`
		Other    string `default:"charset=utf8mb4&parseTime=True"`
	}
	Redis struct {
		Host     string `default:"127.0.0.1"`
		Port     string `default:"6379"`
		Password string `default:"zxlssss"`
	}
}{}
