package main

import (
	"context"
	"log"

	"github.com/alexsuriano/grpcStudy/pb"
	"google.golang.org/grpc"
)

func main() {

	serverGRPC := "localhost:5000"

	conn, err := grpc.Dial(serverGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello World",
	}

	res, err := client.RequestMessage(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}
