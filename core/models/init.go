package models

import (
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

var Engine = Init()
var RDB = InitRedis()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8mb4")
	if err != nil {
		log.Printf("Xorm Error:%v", err)
		return nil
	}
	return engine
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
