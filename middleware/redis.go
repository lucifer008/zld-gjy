package middleware

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"zld-jy/config"
)

var ctx = context.Background()

//const redis_address = "124.223.63.156:6379"
//const redis_password = "zxl_1983131313"

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host + ":" + config.Config.Redis.Port,
		Password: config.Config.Redis.Password, // no password set
		DB:       0,                            // use default DB
	})
	log.Println(">>>>>>>Redis初始化成功>>>>>>>>>")
}
func Set(key string, val interface{}) {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}
}
func Get(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
