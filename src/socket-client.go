package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	conn,err := net.Dial("tcp",":8848")
	if err != nil {
		fmt.Println("net.dial err ",err)
		return
	}
	fmt.Println("Client 与 Server建立连接成功!")

	//向服务器发送数据
	sendData := []byte("hello world")
	for {
		cnt,err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn.write err ",err)
			return
		}
		fmt.Println("Client ===> Server,长度 ",cnt,"，数据",string(sendData))
		//接收服务器返回的数据
		//创建buf,用于接收服务器返回的数据
		buf := make([]byte,1024)
		cnt,err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.read err ",err)
			return
		}
		fmt.Println("Client <=== Server,长度 ",cnt,"，数据",string(buf))

		time.Sleep(1*time.Second)
	}
	conn.Close()

}
