package main

import (
	"fmt"
	"time"
)

func main (){
	//var chan1,chan2 chan int
	chan1 := make(chan int)
	chan2 := make(chan int)

	//启动一个go程，负责监听两个channel
	go func(){
		for {
			select {
			case data1:= <-chan1:
				fmt.Println("从channel1读取数据",data1)
			case data2:= <-chan2:
				fmt.Println("从channel2读取数据",data2)
			default:
				fmt.Println("select 被阻塞")
				time.Sleep(time.Second)
			}
		}

	}()
	//启动go1 写chan1
	go func() {
		for i:=0;i<10;i++{
			chan1 <-i
			time.Sleep(1*time.Second/2)
		}
	}()
	//启动go2 写chan2
	go func() {
		for i:=0;i<10;i++{
			chan2 <-i
			time.Sleep(1*time.Second)
		}
	}()
	for {
		fmt.Println("over")
		time.Sleep(5*time.Second)
	}
}
