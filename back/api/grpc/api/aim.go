package api

import (
	"context"
	"log"

	db "github.com/j-kato732/aimo/db"
	pb "github.com/j-kato732/aimo/proto"
)

const (
	db_path = "./db/aimo.db"
)

var (
	aim_orm          pb.AimModelORM
	response_status  int64
	response_message string
)

type getAimoService struct {
	pb.UnimplementedAimoServer
}

func (s *getAimoService) PostAim(ctx context.Context, aim_model *pb.AimModel) (*pb.PostAimResponse, error) {
	result, err := db.PostAim(ctx, aim_model)
	if err != nil {
		log.Println(err)
	} else {
		response_status = 1
		response_message = ""
	}

	return &pb.PostAimResponse{
		Response: &pb.DefaultResponse{
			Status:  response_status,
			Message: response_message,
		},
		Result: &pb.PostAimResponse_PostAimResult{
			Id: 1,
		},
	}, nil

}
