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

func (*AimoService) GetUser(ctx context.Context, request *pb.UserModel) (*pb.GetUserResponse, error) {
	var response *pb.GetUserResponse = new(pb.GetUserResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetUser(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNoContent) {
			response.Response = newDefaultResponse(normal_code, err.Error())
			return response, nil
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetUserResponse_GetUserResult)
	response.Result.User = result

	return response, nil
}

func (*AimoService) PostUser(ctx context.Context, request *pb.UserModel) (*pb.PostDefaultResponse, error) {
	// request null validate
	// request format valid
	// post実行
	result, err := db.PostUser(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrRecordExist) {
			return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		}
		log.Println(err)
		return nil, err
	}
	// response組み立て
	var response *pb.PostDefaultResponse = new(pb.PostDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.PostDefaultResponse_PostResult)
	response.Result.Id = result

	return response, nil
}

func (*AimoService) PutUser(ctx context.Context, request *pb.UserModel) (*pb.PutDefaultResponse, error) {
	// request null valid
	// request format valid
	// put実行
	err := db.PutUser(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}
