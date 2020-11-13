package main

import (
	"RiYu/mongodb"
	"RiYu/server/log"
	"RiYu/server/service"
	"net/http"
)

func main() {
	// 启动日志服务
	if err := log.Start(); err != nil {
		panic(err)
	}
	// 启动数据库服务
	client, err := mongodb.Start()
	if err != nil {
		panic(err)
	}
	if err = service.Init(client); err != nil {
		panic(err)
	}
	// 启动业务服务
	panic(http.ListenAndServe(":7256", service.Router()))
}
