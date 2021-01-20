package main

import (
	"log"

	"github.com/linqiurong2021/todo/conf"
	"github.com/linqiurong2021/todo/server"
)

func main() {
	//
	err := conf.InitConfig("../../conf/conf.ini")
	if err != nil {
		log.Fatal("init config error: ", err)
	}
	// 获取服务启用端口号
	port := conf.Conf.AppConfig.Port
	// 创建服务并启用
	err = server.NewTodoServer(port)
	if err != nil {
		log.Fatal("error:", err)
	}
}
