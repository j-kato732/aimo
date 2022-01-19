package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "github.com/j-kato732/aimo/db"
	errdetails "github.com/j-kato732/aimo/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/j-kato732/aimo/proto"
)

func (s *AimoService) GetAim(ctx context.Context, request *pb.AimModel) (*pb.GetAimResponse, error) {
	var response *pb.GetAimResponse = new(pb.GetAimResponse)

	// validate
	err := request.Validate()
	if err != nil {
		log.Println(err.Error())
		_ = errdetails.AddErrorDetail(ctx, err)
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	// 必須パラメータチェック
	params := map[string]interface{}{
		"period":  request.GetPeriod(),
		"user_id": request.GetUserId(),
	}
	err = requestParamCheck(params)
	if err != nil {
		log.Println(err.Error())
		message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
		response.Response = newDefaultResponse(invalid_param_format, message)
		return response, nil
	}

	// Get実行
	result, err := db.GetAim(ctx, request)
	if err != nil {
		log.Println(err.Error())
		response.Response = newDefaultResponse(255, err.Error())
		return response, nil
	}

	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetAimResponse_GetAimResult)
	response.Result.Aim = result

	return response, nil
}

func (s *AimoService) PostAim(ctx context.Context, post_request_aim *pb.AimModel) (*pb.PostAimResponse, error) {
	var response *pb.PostAimResponse = new(pb.PostAimResponse)

	err := post_request_aim.Validate()
	if err != nil {
		log.Println(err.Error())
		_ = errdetails.AddErrorDetail(ctx, err)
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	// 必須パラメータチェック
	params := map[string]interface{}{
		"period":  post_request_aim.GetPeriod(),
		"user_id": post_request_aim.GetUserId(),
	}
	err = requestParamCheck(params)
	if err != nil {
		message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
		response.Response = newDefaultResponse(invalid_param_format, message)
		return response, nil
	}

	// post実行
	result, err := db.PostAim(ctx, post_request_aim)
	if err != nil {
		if errors.Is(err, errdetails.ErrRecordExist) {
			return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		}
		log.Println(err)
		return nil, err
	} else {
		response_status = 1
		response_message = ""
	}

	return &pb.PostAimResponse{
		Response: &pb.DefaultResponse{
			Status:  response_status,
			Message: response_message,
		},
		Result: result,
	}, nil
}

func (s *AimoService) PutAim(ctx context.Context, request *pb.AimModel) (*pb.PutAimResponse, error) {
	var response *pb.PutAimResponse = new(pb.PutAimResponse)

	params := map[string]interface{}{
		"id":      request.GetId(),
		"period":  request.GetPeriod(),
		"user_id": request.GetUserId(),
	}
	// 必須パラメータチェック
	err := requestParamCheck(params)
	if err != nil {
		message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
		log.Println(buildInvalidParamsMessage(message))
		response.Response = newDefaultResponse(insufficient_param_error, message)
		return response, nil
	}

	err = db.PutAim(ctx, request)
	if err != nil {
		log.Println(err.Error())
		response.Response = newDefaultResponse(255, err.Error())
		return response, nil
	}

	response.Response = newDefaultResponse(normal_code, "")
	return response, nil
}
