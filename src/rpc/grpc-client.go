package main

import (
	"context"
	"fmt"
	"github.com/justcy/uc/src/rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	//1.链接grpc
	grpcConn,err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc dial err",err)
	}
	//2.初始化客户端
	grpcClient := pb.NewSayNameClient(grpcConn)
	var teacher pb.Teacher
	teacher.Name = "itcast"
	teacher.Age = 18
	//3.调用远程函数
	grpcClient.SayHello(context.TODO(),&teacher)
}
