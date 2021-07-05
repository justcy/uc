package main

import "fmt"

func main()  {
	p1 := testPtr1()
	fmt.Println("p1",p1)
	fmt.Println("*p1",*p1)
}
func testPtr1() *string{
	name :="zhangsan"
	p0 := &name
	fmt.Println("*p0",*p0)
	city:="深圳"
	ptr:=&city
	return ptr
}