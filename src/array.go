package main

import "fmt"

func main()  {
	//1-定义，敌营一个具有10个数字的数组
	//C语言：int nums[10] = {1,2,3,4}
	//go定义:
	//nums:=[10]int{1,2,3,4} (常用方式)
	//var nums =[10]int{1,2,3,4}
	//var nums [10]int = [10]int{1,2,3,4}
	nums:=[10]int{1,2,3,4}

	//2-遍历，方式1
	for i:=0;i<len(nums);i++{
		fmt.Println("i:",i,",j:",nums[i])
	}
	//方式2 for range ==>python 支持
	for key,value := range nums{
		//key=0,value=1=>num[0]
		value +=1
		//value全程是一个临时变量，不断被重新复制，修改不会修改原始数组
		fmt.Println("key：",key,",value:",value,",num:",nums[0])
	}
	//在golang中，如果想忽略一个值，可以使用 _
	//如果两个都忽略，那么不能使用:= 而应该直接使用=
	//:=左边必须要有新变量
	for _,value := range nums{
		//key=0,value=1=>num[0]
		value +=1
		//value全程是一个临时变量，不断被重新复制，修改不会修改原始数组
		fmt.Println("_忽略 key value:",value)
	}
	//不定长数组定义
	//3-使用make进行数组创建

}
