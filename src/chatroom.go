package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type User struct {
	//名字
	name string
	//唯一Id
	id string
	//管道
	msg chan string
}

//创建一个mao保存全局user
var allUsers = make(map[string]User)
//定义message通道，用来接收任何发来的消息
var message = make(chan string, 10)

func main() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net listener err", err)
		return
	}
	//启动全局广播go程,写消息给每个用户
	go broadcast()

	fmt.Println("服务器启动成功", err)
	for {
		fmt.Println("监听中…………")
		//需求: server 可接受多个连接  ===>主go程负责监听，子go程负责处理
		//每个链接可以接收处理多轮数据请求

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net listen Accept err ", err)
		}
		fmt.Println("连接建立")
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	fmt.Println("启动业务")
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("clientAddr:", clientAddr)
	//创建user
	newUser := User{
		name: clientAddr,            //id 不可修改
		id:   clientAddr,            //名称，可以修改
		msg:  make(chan string, 10), //注意make分配空间，否则不能写数据
	}
	//添加user 到map
	allUsers[newUser.id] = newUser

	//全局的退出信号，统一处理客户端退出
	var isQuit, restTimer = make(chan bool, 10), make(chan bool, 10)
	go watch(&newUser, conn, isQuit, restTimer)

	//启动go程，负责将msg消息返回给客户端
	go writeBackeToClient(&newUser, conn)
	//向message写入当前用户上线，用于通知所有人(广播)
	loginInfo := fmt.Sprintf("%s:%s===>上线了！", newUser.id, newUser.name)
	message <- loginInfo

	//TODO 以后实现
	for {
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)

		if cnt == 0 {
			fmt.Println("客户端ctrl+c，准备退出！")
			//发送退出信号，统一做退出处理，使用新的管道，来进行信号处理
			isQuit <- true
		}
		if err != nil {
			fmt.Println("conn read err ", err)
			return
		}
		fmt.Println("Client ===> Server,长度 ", cnt, "，数据", string(buf[:cnt+1]), ",cnt:", cnt)

		//业务逻辑处理开始
		//1.查询当前所有用户
		//a.先判断接受的字符串是否是who==>长度&&字符串
		userInput := string(buf[:cnt-1]) //最后一个是回车 去掉
		if len(userInput) == 3 && userInput == "\\who" {
			fmt.Println("用户即将查询所有用户信息!")
			var userInfos []string
			//b.遍历allUsers,拼接字符串，返回给客户端
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("userid:%s,username:%s", user.id, user.name)
				userInfos = append(userInfos, userInfo)
			}
			//最终拼成一个字符串
			r := strings.Join(userInfos, "\n")
			//strings.Split() 分割字符串
			//返回给查询的结果返回给客户端
			newUser.msg <- r
		} else if len(userInput) > 8 && userInput[:6] == "\\rename" {
			newUser.name = strings.Split(userInput, "|")[1]
			allUsers[newUser.id] = newUser
			newUser.msg <- "rename successfully"

		} else {
			message <- userInput
		}
		restTimer <- true

		//业务逻辑处理结束

		////将数据转大写
		//upperData := strings.ToUpper(string(buf))
		//cnt,err = conn.Write([]byte(upperData))
		//fmt.Println("Client <=== Server,长度 ",cnt,"，数据",upperData)
	}
	//关闭连接
	//_ = conn.Close()
}
func broadcast() {
	fmt.Println("广播go程启动成功...")
	defer fmt.Println("broadcast 退出...")
	for {
		fmt.Println("broadcast 监听message中...")
		info := <-message
		for _, value := range allUsers {
			//如果msg是非缓冲，会阻塞
			value.msg <- info
		}
	}

}
func writeBackeToClient(user *User, conn net.Conn) {
	fmt.Printf("user:%s的go程正在监听自己的msg管道:\n", user.name)
	for data := range user.msg {
		fmt.Printf("user: %s，写回给客户端数据位：%s\n", user.name, data)

		_, _ = conn.Write([]byte(data))

	}
}

//启动一个go程，负责监听退出信号，运行清理工作，
func watch(user *User, conn net.Conn, isQuit, resetTime <-chan bool) {
	fmt.Println("watch go 程 监听中")
	defer fmt.Println("watch go 程退出")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s exit already\n ", user.name)
			fmt.Println("删除当前用户", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			logoutInfo := fmt.Sprintf("%s exit already\n ", user.name)
			fmt.Println("删除当前用户", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo
			conn.Close()
			return
		case <-resetTime:
			fmt.Printf("连接%s重置计数器!\n", user.name)
			return

		}
	}
}
