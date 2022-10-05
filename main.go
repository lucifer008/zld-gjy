package main

import (
	"zld-jy/routers"
)

func init() {

	//todo 单元测试时注释以下代码
	//conf := flag.String("con", "app.yml", "配置文件")
	//flag.Parse()
	//config.Load(*conf)
	//log.Println(">>>>>>>>>>配置信息>>>>>>>>>>>>", config.Config)
	//base.InitDB()
	//middleware.InitRedis()
}
func main() {
	routers.Run()
}
