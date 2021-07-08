package main

import (
	"fmt"
	"net"
	//"net/rpc"
	"net/rpc/jsonrpc"
)

type world struct{
}
func (this *world) HelloWorld(name string,resp *string) error{
	*resp = name +" 你好!"
	return nil
}
func main()  {
	//1.注册RPC服务，绑定对象方法
	RegisterService(new(world))
	//err := rpc.RegisterName("hello", new(world))
	//if err !=nil {
	//	fmt.Println("注册RPC服务失败",err)
	//	return
	//}
	//2.设置监听
	listener, err := net.Listen("tcp", ":8800")

	if err !=nil {
		fmt.Println("net.listen err",err)
		return
	}
	defer  listener.Close()
	fmt.Println("开始监听")
	//3.建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept() err",err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功...")
	//4.绑定服务
	jsonrpc.ServeConn(conn)
}
