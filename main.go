package main

import (
	"flag"
	"log"
	"zld-jy/config"
	"zld-jy/da/base"
	"zld-jy/routers"
)

func init() {
	conf := flag.String("c", "app.yml", "配置文件")
	flag.Parse()
	config.Load(*conf)
	log.Println(">>>>>>>>>>配置信息>>>>>>>>>>>>", config.Config)
	base.InitDB()

	//log.Println(">>>>>>>>>>redis 初始化成功!>>>>>>>>>>>>", config.Config)
}
func main() {
	routers.Run()
}
