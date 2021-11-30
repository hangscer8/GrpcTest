package main

import (
	hello_grpc "GrpcTest/pb"
	"context"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer //1.为未实现的结构体实现方法 SayHi并未具体实现，此时需要将其具体实现
}

func (receiver *server) SayHi(ctx context.Context, req *hello_grpc.Req) (*hello_grpc.Res, error) { //2. 挂载方法
	//一个简单的实现
	return &hello_grpc.Res{Message: req.GetMessage() + ",哈哈哈okokok，返回给客户端"}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":15050")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{}) //注册服务 如果有多个服务，则写多个
	s.Serve(lis)                                     //八股文固定组合写法
}
