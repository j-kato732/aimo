package service

import (
	"context"
	"errors"
	"log"

	db "github.com/j-kato732/aimo/db"
	errdetails "github.com/j-kato732/aimo/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/j-kato732/aimo/proto"
)

/*
/policy
*/
func (*AimoService) GetPolicy(ctx context.Context, request *pb.PolicyModel) (*pb.GetPolicyResponse, error) {
	var response *pb.GetPolicyResponse = new(pb.GetPolicyResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetPolicy(ctx, request)
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
	response.Result = new(pb.GetPolicyResponse_GetPolicyResult)
	response.Result.Policy = result

	return response, nil
}

func (*AimoService) PostPolicy(ctx context.Context, request *pb.PolicyModel) (*pb.PostDefaultResponse, error) {
	// request null validate
	// request format valid
	// post実行
	result, err := db.PostPolicy(ctx, request)
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

func (*AimoService) PutPolicy(ctx context.Context, request *pb.PolicyModel) (*pb.PutDefaultResponse, error) {
	// request null valid
	// request format valid
	// put実行
	err := db.PutPolicy(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}
