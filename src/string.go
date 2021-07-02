package main

import "fmt"

func main(){
	name :="zhangsan"

	//需要输出换行
	usage :=`./a.out <option>
		-h help
	   -a xxx`
	fmt.Println("name :",name)
	fmt.Println("usage :",usage)
	//2.长度访问
	//C++: name.length
	//GO:string 没有.length方法，可以使用自由函数len()
	//len：很常用
	l1 := len(name)
	fmt.Println("l1:",l1)

	//不需要加()
	for i:=0;i<len(name);i++{
		fmt.Printf("i:%d,v:%c\n",i,name[i])
	}
	//3-拼接
	i,j := "hello","world"
	fmt.Println("i+j=",i+j)

	//使用const修复为常量，不能修改
	const address  = "beijing"//在编译期就确定了类型，是预处理，所以不需要推导
	const test  = 100
	//address ="上海"
	fmt.Println("address:",address)
}
