package service

import (
	"context"
	"errors"
	"log"

	db "github.com/j-kato732/aimo/db"
	errdetails "github.com/j-kato732/aimo/errors"
	pb "github.com/j-kato732/aimo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*AimoService) GetPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (*pb.GetPersonalEvaResponse, error) {
	var response *pb.GetPersonalEvaResponse = new(pb.GetPersonalEvaResponse)

	// request null validate
	// request format validate
	// get呼び出し
	result, err := db.GetPersonalEva(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetPersonalEvaResponse_GetPersonalEvaResult)
	response.Result.PersonalEva = result

	return response, nil
}

func (*AimoService) PostPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (*pb.PostDefaultResponse, error) {
	var response *pb.PostDefaultResponse = new(pb.PostDefaultResponse)
	// request null validate
	// request format validate
	// post呼び出し
	result, err := db.PostPersonalEva(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrRecordExist) {
			return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.PostDefaultResponse_PostResult)
	response.Result.Id = result

	return response, err
}

func (*AimoService) PutPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (*pb.PutDefaultResponse, error) {
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)

	// request null validate
	// request format validate
	// put実行
	err := db.PutPersonalEva(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}
