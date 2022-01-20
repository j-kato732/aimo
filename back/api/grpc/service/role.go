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

func (*AimoService) GetRole(ctx context.Context, request *pb.RoleModel) (*pb.GetRoleResponse, error) {
	var response *pb.GetRoleResponse = new(pb.GetRoleResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetRole(ctx, request)
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
	response.Result = new(pb.GetRoleResponse_GetRoleResult)
	response.Result.Role = result

	return response, nil
}

func (*AimoService) PostRole(ctx context.Context, request *pb.RoleModel) (*pb.PostDefaultResponse, error) {
	// request null validate
	// request format valid
	// post実行
	result, err := db.PostRole(ctx, request)
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

func (*AimoService) PutRole(ctx context.Context, request *pb.RoleModel) (*pb.PutDefaultResponse, error) {
	// request null valid
	// request format valid
	// put実行
	err := db.PutRole(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}
