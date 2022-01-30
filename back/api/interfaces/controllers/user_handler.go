package controllers

import (
	pb "github.com/j-kato732/aimo/proto"
	"github.com/j-kato732/aimo/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewUserServerGrpc(gserver *grpc.Server, userUcase usecase.UserInteractor) {
	userServer := &server{
		usecase: userUcase,
	}

	pb.RegisterAimoServer(gserver, userServer)
	reflection.Register(gserver)
}

type server struct {
	usecase usecase.UserInteractor
}
