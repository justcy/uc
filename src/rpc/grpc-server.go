package main

import (
	"context"
	"fmt"
	"github.com/justcy/uc/src/rpc/pb"
	"google.golang.org/grpc"
	"net"
)

type Children struct {

}
func (this *Children) SayHello(ctx context.Context,t *pb.Teacher)(*pb.Teacher,error){
	t.Name += "is Sleeping"
	return t,nil
}

func main() {
	//1.链接grpc
	grpcServer := grpc.NewServer()
	//2.注册服务
	pb.RegisterSayNameServer(grpcServer,new(Children))
	//3.设置监听 指定IP port
	listener, err := net.Listen("tcp","127.0.0.1:8800")
	if err != nil {
		fmt.Println("listener err ",err)
		return
	}
	defer listener.Close()
	//4.启动服务
	grpcServer.Serve(listener)

}
