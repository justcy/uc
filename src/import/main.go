package main

import (
	"fmt"
	"github.com/justcy/uc/src/import/sub"
	//SUB "github.com/justcy/uc/src/import/sub"//SUB是自定包名
	//. "github.com/justcy/uc/src/import/sub"//.代表用户在调用这个包里面的函数时，不需要使用包名.的方式，不建议使用
) //sub是文件名，同时也是包名

func main()  {
	res:=sub.Sub(10,5)
	fmt.Println("sub.Sub(10,5) = ",res)

	//如果包里面的函数想要对外提供访问,那么首字母大写:public
	//大写：public
	//小写:private,只有相同包名的文件
	//add.Add(1,2)
}

