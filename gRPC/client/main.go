package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "learn_Go/gRPC/Proto"
	"log"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 1 连接服务端， 调用 conn = grpc.Dial("localhost:50051", grpc.WithInsecure())
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2 创建 GreeterClient(conn)，把连接成功后返回的 conn 传入
	c := pb.NewGreeterClient(conn)

	// 3 调用具体的方法 SayHello()
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//我们创建并填充一个 HelloRequest 发送给服务。
	//我们用请求调用存根的 SayHello()，如果 RPC 成功，会得到一个填充的 HelloReply ，从其中我们可以获得 greeting。
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
