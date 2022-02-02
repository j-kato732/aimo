package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	pb "github.com/j-kato732/aimo/proto"
	"go.uber.org/zap"

	"github.com/j-kato732/aimo/grpc/interceptor"
	"github.com/j-kato732/aimo/grpc/service"
)

const (
	port = ":8080"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Println(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v\n", err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_auth.UnaryServerInterceptor(interceptor.AuthFuncHandler),
		)),
	)
	// 実行したい処理をserverに登録する
	pb.RegisterAimoServer(server, &service.AimoService{})
	log.Printf("server listening at %v\n", lis.Addr())
	if err != nil {
		log.Fatalf("faild to serve: %v\n", err)
	}
	server.Serve(lis)
}
