package main

import "fmt"

//1.函数返回值在参数列表之后
//2.如果有多个返回值，需要用圆括号包裹，多个参数之间用，分隔
func test2(a int ,b int ,c string)(int, string ,bool){
	return a+b ,c,true
}

func test3(a int ,b int ,c string)(res int, str string ,bl bool){
	var i,j,k int
	i = 1
	j = 2
	k = 3
	fmt.Println(i,j,k)
	//直接使用返回值的变量名字参与计算
	res = a+b
	str = c
	bl =true
	//当返回值有名字的时候，可以直接写return
	return
	//return a+b ,c,true
}
//如果返回值只有一个参数，并且没有名字，那么不需要加括号
func test4()int{
	return 10
}
func main(){
	v1,s1,_:=test2(10,20,"hello")
	fmt.Println("v1:",v1,",S1:",s1)

	v3,s3,_:=test3(20,30,"world")
	fmt.Println("v3:",v3,",S3:",s3)
}
