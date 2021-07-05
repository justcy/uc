package main

import (
	"fmt"
	"os"
)

//从命令行传入参数，在switch中处理
func main()  {
	//C: argc ** argv
	//go : os.Args ==>直接可以获取命令输入，是一个字符串切片
	cmds := os.Args
	//os.Args[0] ==>程序名字
	//os.Args[1] ==>第一个参数，以此类推

	for key,value := range cmds{
		fmt.Println("key:",key,",cmd:",value,",cmds len ",len(cmds))
	}
	if len(cmds) < 1{
		fmt.Println("请正确输入参数!")
	}
	//go的switch默认加上了break，不需要手动处理
	//如果想向下穿透的话，需要加上关键词 fallthrough
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default called")
	}
}
