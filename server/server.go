package main

import(
	"log"
	"../service"
	"./calculate"
)

func main(){

	log.Println("Starting service...")
	//注册服务
	service.Register(new(calculate.Pairs))
	//启动服务器
	service.Start("tcp", ":8081")
}

