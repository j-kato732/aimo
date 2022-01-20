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

func (*AimoService) GetAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.GetAchievementMeanResponse, error) {
	var response *pb.GetAchievementMeanResponse = new(pb.GetAchievementMeanResponse)
	// request null validate
	// request format validate
	// get実行
	result, err := db.GetAchievementMean(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNoContent) {
			response.Response = newDefaultResponse(normal_code, err.Error())
			return response, nil
		}
		log.Println(err.Error())
		return nil, err
	}

	// response生成
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetAchievementMeanResponse_GetAchievementMeanResult)
	response.Result.AchievementMean = result

	return response, nil
}

func (*AimoService) PostAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.PostAchievementMeanResponse, error) {
	var response *pb.PostAchievementMeanResponse = new(pb.PostAchievementMeanResponse)
	// request nul validate
	// request format validate
	// post実行
	result, err := db.PostAchievementMean(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrRecordExist) {
			return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		}
		log.Println(err.Error())
		return nil, err
	}
	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.PostAchievementMeanResponse_PostAchievementMeansResult)
	response.Result.Id = result

	return response, nil
}

func (*AimoService) PutAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.PutDefaultResponse, error) {
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	// request null validate
	// request format validate
	// put実行
	err := db.PutAchievementMean(ctx, request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}
