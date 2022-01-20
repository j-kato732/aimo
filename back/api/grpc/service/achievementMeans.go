package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "github.com/j-kato732/aimo/db"
	errdetails "github.com/j-kato732/aimo/errors"
	pb "github.com/j-kato732/aimo/proto"
)

func (s *AimoService) GetAchievementMeans(ctx context.Context, request *pb.AchievementMeanModel) (*pb.GetAchievementMeansResponse, error) {
	// aim_number := request.GetAimNumber()
	// achievement_mean_number := request.GetAchievementMeanNumber()

	var response *pb.GetAchievementMeansResponse = new(pb.GetAchievementMeansResponse)
	// fmt.Println((*request).Period)

	params := map[string]interface{}{
		"period":     request.GetPeriod(),
		"user_id":    request.GetUserId(),
		"aim_number": request.GetAimNumber(),
	}
	// 必須パラメータチェック
	err := requestParamCheck(params)
	if err != nil {
		message := fmt.Sprintf("%+v", request)
		log.Println(buildInvalidParamsMessage(message))
		response.Response = newDefaultResponse(insufficient_param_error, buildInvalidParamsMessage(message))
		return response, nil
	}

	// 具体的達成手段取得
	result, err := db.GetAchievementMeans(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNoContent) {
			response.Response = newDefaultResponse(normal_code, err.Error())
			return response, nil
		}
		log.Println(err.Error())
		response.Response = newDefaultResponse(255, err.Error())
	}

	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetAchievementMeansResponse_GetAchievementMeansResult)
	response.Result.AchievementMeans = result

	// log.Printf("%+v", request)
	return response, nil
}
