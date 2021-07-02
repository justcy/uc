package main

import "fmt"

func main(){
	//变量定义： var
	//常量定义: const
	//01-先定义变量，再赋值 var 变量名 数据类型
	var name string
	var age int
	name = "zhangsan"
	age = 20
	fmt.Println(name)
	fmt.Printf("name is %s,%d",name,age)

	//02-定义直接赋值
	var gender = "man"
	fmt.Println("gender ",gender)

	//03-定义直接赋值，使用推导 最常用
	address := "beijing"
	fmt.Println("gender ",address)

	//04-平行赋值
	i,j := 10,20
	fmt.Println("变换前 i:",i,",j",j)
	i,j = j,i
	fmt.Println("变换后 i:",i,",j",j)
}
