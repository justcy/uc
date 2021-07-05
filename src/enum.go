package main

import "fmt"

const (
	MON= iota
	TUE
	WED
	THU
	FRI
	SAT
	SUN
	M,N = iota,iota //const属于预编译期赋值，不需要:=自动推导
)
const (
	JAN= iota
	FER
	MAR
)

//1.iota是常量组计数器
//2.iota 从0 开始每次换行+1
//3.常量租有个特点，如果不赋值，默认与上一行相同
//4.如果同意行出现iota，那么两个iota的值是相同的
//5.每个常量租的iota是独立的，如果遇到const iota会重新清0
func main()  {
	fmt.Println("打印周:")
	fmt.Println(MON)
	fmt.Println(TUE)
	fmt.Println(WED)
	fmt.Println(THU)
	fmt.Println(FRI)
	fmt.Println(SAT)
	fmt.Println(SUN)

	fmt.Println("打印月:")
	fmt.Println(JAN)
	fmt.Println(FER)
	fmt.Println(MAR)
}
