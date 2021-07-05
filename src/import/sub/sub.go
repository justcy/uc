package sub
//package utils  //不允许出现多个包名

func Sub(a,b int) int{
	test4()//由于test4与sub.go在同一个包下面，所以可以使用，且不需要sub.形式
	return a-b
}
