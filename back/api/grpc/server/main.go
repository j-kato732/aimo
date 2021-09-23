package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	db "github.com/j-kato732/aimo/db"
	pb "github.com/j-kato732/aimo/proto"
)

const (
	port = ":8080"

	normal_code = 1

	invalid_param_error_code = 10
	invalid_param_format     = 11

	other_error_code = 255

	period_type_error_message   = "period must be of type string"
	period_length_error_message = "period must be of length 6 (example: 202105)"
)

var (
	aim_orm          pb.AimModelORM
	response_status  int64
	response_message string

	post_aim_response pb.PostAimResponse
)

type getAimoService struct {
	pb.UnimplementedAimoServer
}

func (s *getAimoService) GetAim(ctx context.Context, get_aim_request *pb.GetAimRequest) (*pb.GetAimResponse, error) {
	return nil, errors.New("OK")
}

func (s *getAimoService) PostAim(ctx context.Context, post_request_aim *pb.AimModel) (*pb.PostAimResponse, error) {
	if post_request_aim.GetPeriod() == "" || post_request_aim.GetUserId() == 0 {
		log.Println(post_request_aim)
		log.Println("invalid param")
		return nil, errors.New("invalid parameter")
	}

	result, err := db.PostAim(ctx, post_request_aim)
	if err != nil {
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

func (s *getAimoService) PostAchievementMeans(ctx context.Context, post_request_achievement_means *pb.PostAchievementMeansRequest) (*pb.PostAchievementMeansResponse, error) {
	// achievementMeans配列を取得
	post_request_achievement_means_model := post_request_achievement_means.GetAchievementMeans()

	var achievement_means_response []*pb.PostAchievementMeanResponse

	// listの数、新規作成
	for _, post_request_achievement_mean := range post_request_achievement_means_model {
		period := post_request_achievement_mean.GetPeriod()
		user_id := post_request_achievement_mean.GetUserId()
		aim_number := post_request_achievement_mean.GetAimNumber()
		achievement_mean_number := post_request_achievement_mean.GetAchievementMeanNumber()

		var achievement_mean *pb.PostAchievementMeanResponse = new(pb.PostAchievementMeanResponse)

		// 必須パラメータ存在チェック
		if len(period) == 0 || user_id == 0 || aim_number == 0 || achievement_mean_number == 0 {
			message := paramsMessageBuild(map[string]string{
				"period":                 period,
				"user_id":                strconv.Itoa(int(user_id)),
				"aim_number":             strconv.Itoa(int(aim_number)),
				"achievement_aim_number": strconv.Itoa(int(achievement_mean_number))})
			log.Println("invalid required params")
			log.Printf(message)
			achievement_mean.Response = newDefaultResponse(255, buildInvalidParamsMessage(message))
			achievement_means_response = append(achievement_means_response, achievement_mean)
			continue
		}

		// period型チェック
		result, err := validatePeriodTypeFormat(period)
		if err != nil {
			log.Println(err.Error())
			achievement_mean.Response = newDefaultResponse(result, err.Error())
			achievement_means_response = append(achievement_means_response, achievement_mean)
			continue
		}

		// period 長さ、日付フォーマットチェック
		result, err = validatePeriodFormat(period)
		if err != nil {
			log.Println(err.Error())
			achievement_mean.Response = newDefaultResponse(result, err.Error())
			achievement_means_response = append(achievement_means_response, achievement_mean)
			continue
		}

		// 新規作成
		result, err = db.PostAchievementMean(ctx, post_request_achievement_mean)
		if err != nil {
			achievement_mean.Response = newDefaultResponse(255, err.Error())
			achievement_mean.Result = &pb.PostAchievementMeanResponse_PostAchievementMeansResult{Id: result}
		} else {
			achievement_mean.Response = newDefaultResponse(1, "")
			achievement_mean.Result = &pb.PostAchievementMeanResponse_PostAchievementMeansResult{Id: result}
		}
		achievement_means_response = append(achievement_means_response, achievement_mean)
	}

	return &pb.PostAchievementMeansResponse{
		Responses: achievement_means_response,
	}, nil
}

func newDefaultResponse(status int64, message string) *pb.DefaultResponse {
	default_response := new(pb.DefaultResponse)
	default_response.Status = status
	default_response.Message = message
	return default_response
}

func paramsMessageBuild(params map[string]string) string {
	var message string

	params_num := len(params)
	cnt := 0
	for key, param := range params {
		cnt++
		if cnt == 1 {
			message = "(" + key + ": " + param + ", "
			continue
		}

		if cnt == params_num {
			message = message + key + ": " + param + ")"
			break
		}

		message = message + key + ": " + param + ", "

	}
	return message
}

func buildInvalidParamsMessage(message string) string {
	return "invalid require params " + message
}

func validatePeriodFormat(period string) (int64, error) {
	// 長さチェック
	if len(period) != 6 {
		return invalid_param_format, errors.New(period_length_error_message)
	}

	// TODO:日付フォーマットチェック

	return 1, nil
}

func validatePeriodTypeFormat(period interface{}) (int64, error) {
	switch period.(type) {
	case string:
		return normal_code, nil
	default:
		return invalid_param_format, errors.New(period_type_error_message)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v\n", err)
	}
	server := grpc.NewServer()
	// 実行したい処理をserverに登録する
	pb.RegisterAimoServer(server, &getAimoService{})
	log.Printf("server listening at %v\n", lis.Addr())
	if err != nil {
		log.Fatalf("faild to serve: %v\n", err)
	}
	server.Serve(lis)
}
