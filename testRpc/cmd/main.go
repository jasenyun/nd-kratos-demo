package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	pb "jasenServer/api/helloworld/v1"
	"log"
)

func main() {

	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("127.0.0.1:9000"),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	// returns error
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Printf("[grpc] SayHello error: %v\n", err)
	}
	log.Printf("[grpc] SayHello success: %v\n", reply.Message)
}
