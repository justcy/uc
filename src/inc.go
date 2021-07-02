package main

import "fmt"

func main()  {
	i:=20
	fmt.Println("i：",i)

	i++
	//++i//不支持这种语法
	//fmt.Println("i：",i++)//错误，不允许与其他代码一起，需单独起一行
	fmt.Println("i：",i)
}