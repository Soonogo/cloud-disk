package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8mb4")
	if err != nil {
		log.Printf("Xorm Error:%v", err)
		return nil
	}
	return engine
}
