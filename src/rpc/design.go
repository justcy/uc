package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//要求服务端在注册RCP对象时，能让编译器检测出，对象是否合法
//创建接口，在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

//调用该方法，需要给I传参数，参数是实现了HelloWorld方法的类对象！
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

//定义类
type MyClient struct {
	c *rpc.Client
}

//由于使用了c调用Call,因此要初始化C
func InitClient(addr string) MyClient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return MyClient{c: conn}
}

//实现函数，原型上参考interface的实现
func (this *MyClient) HelloWorld(a string, b *string) error {
	//1.参数1 参照上面的interface  RegisterName而来，a:传入参数 b:传出参数
	return this.c.Call("hello.HelloWorld", a, b)
}