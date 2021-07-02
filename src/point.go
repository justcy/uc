package main

import "fmt"

func main()  {
	//go也有指针
	//结构体成员调用时：C语言ptr->name go:ptr.name
	//go语言在使用指针时，会使用内部垃圾回收机制，(gc,garbage collector),不需要手动释放指针
	//C语言不允许返回栈上指针，go语言可以返回栈上指针，程序会在编译的时候确定变量分配的位置。
	//编译时如果发现有必要就会将变量分配到堆上，内存逃逸。

	//01-取自变量
	name := "zhangsan"
	ptr := &name
	fmt.Println("name:",name)
	fmt.Println("name ptr:",ptr)

	//02-使用new关键词定义
	name2ptr := new(string)
	*name2ptr = "wangwu"
	fmt.Println("name2:",*name2ptr)
	fmt.Println("name2 ptr:",name2ptr)

	//可以返回栈上指针，编译器在编译程序时，会自动判断这段代码，将city变量分配在堆上，内存逃逸
	res := testpPtr()
	fmt.Println("res city:",*res,",address ",res)

	//空指针，在C语言：null C++:nullptr,go:nil
	//if 不带() ,即使一行代码也必须带 {}
	if res== nil{
		fmt.Println("res 是空,nil")
	}else{
		fmt.Println("res 是非空,nil")
	}
}
//定义一个函数，返回一个string类型的指针，go语言返回写在参数列表后面
func testpPtr() *string{
	city := "武汉"
	ptr := &city
	return  ptr
}