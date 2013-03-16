package service

import(
	"net/rpc"
	"net/http"
	"net"
	"log"
)

//注册服务
func Register(v interface{}){

	rpc.Register(v)
}

//启动服务器
func Start(protocal string, port string){

	rpc.HandleHTTP()

	l, e:= net.Listen(protocal, port)

	if e != nil{
		log.Panic("listen error: ", e)
	}

	http.Serve(l, nil)
}

//根据地址调用服务，传入参数，获得结果
func Call(address string, service string, args interface{}, reply interface{}) error{

	client , err := rpc.DialHTTP("tcp", address)

	if err != nil{
		log.Fatal("dialing ", err)
	}

	err = client.Call(service, args, reply)

	return err
}

