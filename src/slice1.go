package main

import "fmt"

func main()  {
	names := [7]string{"北京","上海","广州","深圳","洛阳","南京","秦皇岛"}
	//基于names创建新的数组
	names1 := [3]string{}
	names1[0] = names[0]
	names1[1] = names[1]
	names1[2] = names[2]
	//切片可以基于一个数组，灵活创建新的数组
	name2 := names[0:3]//左闭右开
	fmt.Println("names2",name2)

	name2[2] = "hello"
	fmt.Println("修改name2之后,name2 ",name2)
	fmt.Println("修改name2之后,names",names)

	//如果想切片完全独立于原数组，可以用copy()函数来完成

	//1.如果从0元素开始截取，那么冒号左边的数字可以省略
	names3 := names[:5]
	fmt.Println("names3 ",names3)
	//2.如果截取到最后的元素，那么冒号右边的数字可省略
	names4 := names[5:]
	fmt.Println("names4 ",names4)
	//1.如果想全部使用，那么冒号左右两边都可以省略
	names5 := names[:]
	fmt.Println("names5 ",names5)
	//4.也可以基于字符串进行切片截取，取字符串的子串
	subl := "Helloworld"[5:7]
	fmt.Println("subl ",subl)

	//5.可以在创建空切片的时候，明确指定切片的容量，这样可提高运行效率
	//创建一个容量为20，当前长度是0的string类型切片
	str2 := make([]string,3,20)//第三个参数不是必须的，默认与长度相同
	fmt.Println("str2: len:",len(str2),",cap:",cap(str2))
	str2[0] ="Hello"
	str2[1] ="world"
	//6.如果想让切片完全独立于原始数组，可使用copy()函数来完成
	nameCopy := make([]string,len(names))
	//func copy(dst,src []Type) int
	//names是一个定长数组，copy函数所需要的参数为切片，所以要使用[:],将数组转为切片
	copy(nameCopy,names[:])
	nameCopy[0]= "香港"
	fmt.Println("nameCopy ",nameCopy)
	fmt.Println("names ",names)

}
