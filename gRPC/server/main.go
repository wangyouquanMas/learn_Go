package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "learn_Go/gRPC/Proto"
	"log"
	"net"
)

// Server 端代码提供 gRpc 服务

// 1 指定需要提供服务的端口，本地未被使用的任意端口都可以，比如 50051。

const (
	port = ":50051"
)

// 3 定义一个 server struct 来实现 proto 文件中 service 部分所声明的所有 RPC 方法。 这步最关键！
// server 实现了 greeter服务接口
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "HelloAgain " + in.Name}, nil
}

func main() {
	// 2 监听端口，调用net.Listen("tcp", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 4 调用 grpc.NewServer() 来创建一个 Server
	s := grpc.NewServer()
	// 5 把struct 的实例注册到 grpc server 上
	pb.RegisterGreeterServer(s, new(server))

	// 6 调用 s.Serve(lis) 提供 gRpc 服务
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// 将greeter服务绑定到端口，然后启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
