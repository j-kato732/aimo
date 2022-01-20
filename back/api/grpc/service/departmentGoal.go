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

func (*AimoService) GetDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (*pb.GetDepartmentGoalResponse, error) {
	var response *pb.GetDepartmentGoalResponse = new(pb.GetDepartmentGoalResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetDepartmentGoal(ctx, request)
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
	response.Result = new(pb.GetDepartmentGoalResponse_GetDepartmentGoalResult)
	response.Result.DepartmentGoal = result

	return response, nil
}

func (*AimoService) PostDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (*pb.PostDefaultResponse, error) {
	// request null validate
	// request format valid
	// post実行
	result, err := db.PostDepartmentGoal(ctx, request)
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

func (*AimoService) PutDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (*pb.PutDefaultResponse, error) {
	// request null valid
	// request format valid
	// put実行
	err := db.PutDepartmentGoal(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}
