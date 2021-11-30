package main

import (
	hello_grpc "GrpcTest/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	//为了演示，先设置忽略不安全的连接
	conn, err := grpc.Dial("127.0.0.1:15050", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := hello_grpc.NewHelloGRPCClient(conn)
	res, _ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "发送的数据123"})
	fmt.Println(res.GetMessage())
	conn.Close()
}
