package main

import (
	"fmt"
	"net/rpc"
)

func main01() {
	//1.用RPC连接服务器
	conn,err := rpc.Dial("tcp","127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err",err)
	}
	defer conn.Close()
	//2.调用远程函数
	var replay string
	err = conn.Call("hello.HelloWorld","李白",&replay)
	if err != nil {
		fmt.Println("Call err",err)
		return
	}
	fmt.Println("调用结果：",replay)
}

func main() {
	//1.用RPC连接服务器
	MyClient := InitClient("127.0.0.1:8800")
	var replay string
	err := MyClient.HelloWorld("张三",&replay)
	//2.调用远程函数
	if err != nil {
		fmt.Println("Call err",err)
		return
	}
	fmt.Println("调用结果：",replay)
}