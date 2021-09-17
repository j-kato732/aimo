package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/j-kato732/aimo/proto"
)

const (
	port = ":8080"
)

type getAimoService struct {
	pb.UnimplementedAimoServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v\n", err)
	}
	server := grpc.NewServer()
	// 実行したい処理をserverに登録する
	pb.RegisterAimoServer(server, &getAimoService{})
	log.Printf("server listening at %v\n", lis.Addr())
	if err != nil {
		log.Fatalf("faild to serve: %v\n", err)
	}
	server.Serve(lis)
}
