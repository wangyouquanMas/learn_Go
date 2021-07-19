package main

import (
	pb "channel/gRPC/Proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)



const(
	address  = "localhost:50051"
	defaultName = "world"
)

func main(){
	// 1 连接服务端， 调用 conn = grpc.Dial("localhost:50051", grpc.WithInsecure())
	 conn , err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil{
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
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}