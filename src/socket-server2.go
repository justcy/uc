package main

import (
	"fmt"
	"net"
	"strings"
)

type Teache struct {
	Name string `json:"_"`//
	Subject string `json:"Subject_nam"`
	Age int `json:"age,string"`
}
func main(){
	//创建监听
	ip:= "127.0.0.1"
	port:= 8848
	address := fmt.Sprintf("%s:%d",ip,port)
	fmt.Println(address)

	listener,err := net.Listen("tcp",address)
	if err != nil {
		fmt.Println("net listen err ",err)
	}
	for {
		fmt.Println("监听中…………")
		//需求: server 可接受多个连接  ===>主go程负责监听，子go程负责处理
		//每个链接可以接收处理多轮数据请求

		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("net listen Accept err ",err)
		}
		fmt.Println("连接建立")
		go handleFunc(conn)
	}
}

func handleFunc(conn net.Conn){
	for {
		buf := make([]byte,1024)
		cnt,err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err ",err)
		}
		fmt.Println("Client ===> Server,长度 ",cnt,"，数据",string(buf))

		//将数据转大写
		upperData := strings.ToUpper(string(buf))
		cnt,err = conn.Write([]byte(upperData))
		fmt.Println("Client <=== Server,长度 ",cnt,"，数据",upperData)
	}
	//关闭连接
	//_ = conn.Close()
}
