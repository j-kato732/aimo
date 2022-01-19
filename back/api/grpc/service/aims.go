package service

import (
	"context"
	"fmt"
	"log"

	db "github.com/j-kato732/aimo/db"
	errdetails "github.com/j-kato732/aimo/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/j-kato732/aimo/proto"
)

func (s *AimoService) GetAims(ctx context.Context, get_aim_request *pb.GetAimsRequest) (*pb.GetAimsResponse, error) {
	var response *pb.GetAimsResponse = new(pb.GetAimsResponse)

	err := get_aim_request.Validate()
	if err != nil {
		log.Println(err.Error())
		_ = errdetails.AddErrorDetail(ctx, err)
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	// 必須パラメータチェック
	params := map[string]interface{}{
		"period":  get_aim_request.GetPeriod(),
		"user_id": get_aim_request.GetUserId(),
	}
	err = requestParamCheck(params)
	if err != nil {
		log.Println(err.Error())
		message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
		response.Response = newDefaultResponse(invalid_param_format, message)
		return response, nil
	}

	// Get実行
	result, err := db.GetAims(ctx, get_aim_request)
	if err != nil {
		log.Println(err.Error())
		response.Response = newDefaultResponse(255, err.Error())
		return response, nil
	}

	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetAimsResult)
	response.Result.Aims = result

	return response, nil
}
