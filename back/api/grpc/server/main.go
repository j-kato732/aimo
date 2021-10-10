package main

import (
	"context"
	"errors"
	"fmt"
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

	insufficient_param_error = 10
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

/*
/aim
*/

func (s *getAimoService) GetAim(ctx context.Context, get_aim_request *pb.GetAimRequest) (*pb.GetAimResponse, error) {
	var response *pb.GetAimResponse = new(pb.GetAimResponse)

	// 必須パラメータチェック
	params := map[string]interface{}{
		"period":  get_aim_request.GetPeriod(),
		"user_id": get_aim_request.GetUserId(),
	}
	err := requestParamCheck(params)
	if err != nil {
		log.Println(err.Error())
		message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
		response.Response = newDefaultResponse(invalid_param_format, message)
		return response, nil
	}

	// Get実行
	result, err := db.GetAim(ctx, get_aim_request)
	if err != nil {
		log.Println(err.Error())
		response.Response = newDefaultResponse(255, err.Error())
		return response, nil
	}

	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetAimResult)
	response.Result.Aim = result

	return response, nil
}

func (s *getAimoService) PostAim(ctx context.Context, post_request_aim *pb.AimModel) (*pb.PostAimResponse, error) {
	var response *pb.PostAimResponse = new(pb.PostAimResponse)

	// 必須パラメータチェック
	params := map[string]interface{}{
		"period":  post_request_aim.GetPeriod(),
		"user_id": post_request_aim.GetUserId(),
	}
	err := requestParamCheck(params)
	if err != nil {
		message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
		response.Response = newDefaultResponse(invalid_param_format, message)
		return response, nil
	}

	// post実行
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

func (s *getAimoService) PutAim(ctx context.Context, request *pb.AimModel) (*pb.PutAimResponse, error) {
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

/*
/achievementMeans
*/

func (s *getAimoService) GetAchievementMeans(ctx context.Context, request *pb.AchievementMeanModel) (*pb.GetAchievementMeansResponse, error) {
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
		log.Println(err.Error())
		response.Response = newDefaultResponse(255, err.Error())
	}

	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetAchievementMeansResponse_GetAchievementMeansResult)
	response.Result.AchievementMeans = result

	// log.Printf("%+v", request)
	return response, nil
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

func (s *getAimoService) PutAchievementMeans(ctx context.Context, request *pb.PutAchievementMeansRequest) (*pb.PutAchievementMeansResponses, error) {
	var responses *pb.PutAchievementMeansResponses = new(pb.PutAchievementMeansResponses)
	for _, achievement_mean := range request.GetAchievementMeans() {
		var response *pb.PutAchievementMeansResponses_PutAchievementMeanResponse = new(pb.PutAchievementMeansResponses_PutAchievementMeanResponse)
		params := map[string]interface{}{
			"id":                      achievement_mean.GetId(),
			"period":                  achievement_mean.GetPeriod(),
			"user_id":                 achievement_mean.GetUserId(),
			"aim_number":              achievement_mean.GetAimNumber(),
			"achievement_mean_number": achievement_mean.GetAchievementMeanNumber(),
		}

		// 必須パラメータチェック
		err := requestParamCheck(params)
		if err != nil {
			// map_str_param := interfaceToStringMap(params)
			// message := paramsMessageBuild(map_str_param)
			message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
			log.Println(buildInvalidParamsMessage(message))
			response.Response = newDefaultResponse(insufficient_param_error, message)
			responses.Responses = append(responses.Responses, response)
			return responses, nil
		}
		err = db.UpdateAchievementMean(ctx, achievement_mean)
		if err != nil {
			log.Println(err.Error())
			response.Response = newDefaultResponse(255, err.Error())
			responses.Responses = append(responses.Responses, response)
			return responses, nil
		}

		response.Response = newDefaultResponse(normal_code, "")
		responses.Responses = append(responses.Responses, response)
	}
	return responses, nil
}

/*
/achievementMean
*/

func (*getAimoService) GetAchievementMean(ctx context.Context, request *pb.GetAchievementMeanRequest) (*pb.GetAchievementMeanResponse, error) {
	var response *pb.GetAchievementMeanResponse = new(pb.GetAchievementMeanResponse)
	// request null validate
	// request format validate
	// get実行
	result, err := db.GetAchievementMean(ctx, request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// response生成
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetAchievementMeanResponse_GetAchievementMeanResult)
	response.Result.AchievementMean = result

	return response, nil
}

func (*getAimoService) PostAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.PostAchievementMeanResponse, error) {
	var response *pb.PostAchievementMeanResponse = new(pb.PostAchievementMeanResponse)
	// request nul validate
	// request format validate
	// post実行
	result, err := db.PostAchievementMean(ctx, request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.PostAchievementMeanResponse_PostAchievementMeansResult)
	response.Result.Id = result

	return response, nil
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

func requestParamCheck(params map[string]interface{}) error {
	log.Println(params)
	for key, param := range params {
		if value, ok := param.(int64); ok {
			if value == 0 {
				message := fmt.Sprintf("request params require %v", key)
				log.Println("Error: " + message)
				return errors.New(message)
			}
		}

		if value, ok := param.(string); ok {
			if len(value) == 0 {
				message := fmt.Sprintf("request params require %v", key)
				log.Println("Error: " + message)
				return errors.New(message)
			}
		}
	}
	return nil
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
