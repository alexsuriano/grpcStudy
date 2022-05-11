package main

import (
	"context"
	"log"
	"net"

	"github.com/alexsuriano/grpcStudy/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("received message: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}

func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := "localhost:5000"
	conn := "tcp"

	listener, err := net.Listen(conn, port)
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("grpc server error: ", err)
	}
}
