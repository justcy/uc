package main

import "fmt"

func main()  {
	//切片:slice ,它的底层也是数组，可动态改变长度
	//定义一个切片,包含多个地名
	//name := [10]string{"北京","上海","广州","深圳"}
	names := []string{"北京","上海","广州","深圳"}
	for i,v := range names{
		fmt.Println("i:",i,"v:",v)
	}
	//1.追加数据
	//name1
	names1 := append(names, "海南")
	fmt.Println("names ",names)
	fmt.Println("names1 ",names1)
	fmt.Println("追加元素前，name长度：",len(names),",容量:",cap(names))
	names = append(names, "海南")
	fmt.Println("names 追加元素后赋值给自己",names)
	fmt.Println("追加元素后，name长度：",len(names),",容量:",cap(names))
	names = append(names, "四川")
	fmt.Println("追加元素后，name长度：",len(names),",容量:",cap(names))
	//2.对于一个切片，不仅有长度的概念len(),还有一个 容量 的概念
	nums := []int{}
	for i:=0;i<16;i++{
		nums = append(nums, i)
		fmt.Println("nums：",len(nums),",容量:",cap(nums))
	}

}
