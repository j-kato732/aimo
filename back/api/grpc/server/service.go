package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "github.com/j-kato732/aimo/db"
	errdetails "github.com/j-kato732/aimo/errors"
	pb "github.com/j-kato732/aimo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
/aims
*/

func (s *getAimoService) GetAims(ctx context.Context, get_aim_request *pb.GetAimsRequest) (*pb.GetAimsResponse, error) {
	var response *pb.GetAimsResponse = new(pb.GetAimsResponse)

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

func (s *getAimoService) GetAim(ctx context.Context, request *pb.AimModel) (*pb.GetAimResponse, error) {
	var response *pb.GetAimResponse = new(pb.GetAimResponse)

	// 必須パラメータチェック
	params := map[string]interface{}{
		"period":  request.GetPeriod(),
		"user_id": request.GetUserId(),
	}
	err := requestParamCheck(params)
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
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
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

// func (s *getAimoService) PostAchievementMeans(ctx context.Context, post_request_achievement_means *pb.PostAchievementMeansRequest) (*pb.PostAchievementMeansResponse, error) {
// 	// achievementMeans配列を取得
// 	post_request_achievement_means_model := post_request_achievement_means.GetAchievementMeans()

// 	var achievement_means_response []*pb.PostAchievementMeanResponse

// 	// listの数、新規作成
// 	for _, post_request_achievement_mean := range post_request_achievement_means_model {
// 		period := post_request_achievement_mean.GetPeriod()
// 		user_id := post_request_achievement_mean.GetUserId()
// 		aim_number := post_request_achievement_mean.GetAimNumber()
// 		achievement_mean_number := post_request_achievement_mean.GetAchievementMeanNumber()

// 		var achievement_mean *pb.PostAchievementMeanResponse = new(pb.PostAchievementMeanResponse)

// 		// 必須パラメータ存在チェック
// 		if len(period) == 0 || user_id == 0 || aim_number == 0 || achievement_mean_number == 0 {
// 			message := paramsMessageBuild(map[string]string{
// 				"period":                 period,
// 				"user_id":                strconv.Itoa(int(user_id)),
// 				"aim_number":             strconv.Itoa(int(aim_number)),
// 				"achievement_aim_number": strconv.Itoa(int(achievement_mean_number))})
// 			log.Println("invalid required params")
// 			log.Printf(message)
// 			achievement_mean.Response = newDefaultResponse(255, buildInvalidParamsMessage(message))
// 			achievement_means_response = append(achievement_means_response, achievement_mean)
// 			continue
// 		}

// 		// period型チェック
// 		result, err := validatePeriodTypeFormat(period)
// 		if err != nil {
// 			log.Println(err.Error())
// 			achievement_mean.Response = newDefaultResponse(result, err.Error())
// 			achievement_means_response = append(achievement_means_response, achievement_mean)
// 			continue
// 		}

// 		// period 長さ、日付フォーマットチェック
// 		result, err = validatePeriodFormat(period)
// 		if err != nil {
// 			log.Println(err.Error())
// 			achievement_mean.Response = newDefaultResponse(result, err.Error())
// 			achievement_means_response = append(achievement_means_response, achievement_mean)
// 			continue
// 		}

// 		// 新規作成
// 		result, err = db.PostAchievementMean(ctx, post_request_achievement_mean)
// 		if err != nil {
// 			achievement_mean.Response = newDefaultResponse(255, err.Error())
// 			achievement_mean.Result = &pb.PostAchievementMeanResponse_PostAchievementMeansResult{Id: result}
// 		} else {
// 			achievement_mean.Response = newDefaultResponse(1, "")
// 			achievement_mean.Result = &pb.PostAchievementMeanResponse_PostAchievementMeansResult{Id: result}
// 		}
// 		achievement_means_response = append(achievement_means_response, achievement_mean)
// 	}

// 	return &pb.PostAchievementMeansResponse{
// 		Responses: achievement_means_response,
// 	}, nil
// }

// func (s *getAimoService) PutAchievementMeans(ctx context.Context, request *pb.PutAchievementMeansRequest) (*pb.PutAchievementMeansResponses, error) {
// 	var responses *pb.PutAchievementMeansResponses = new(pb.PutAchievementMeansResponses)
// 	for _, achievement_mean := range request.GetAchievementMeans() {
// 		var response *pb.PutAchievementMeansResponses_PutAchievementMeanResponse = new(pb.PutAchievementMeansResponses_PutAchievementMeanResponse)
// 		params := map[string]interface{}{
// 			"id":                      achievement_mean.GetId(),
// 			"period":                  achievement_mean.GetPeriod(),
// 			"user_id":                 achievement_mean.GetUserId(),
// 			"aim_number":              achievement_mean.GetAimNumber(),
// 			"achievement_mean_number": achievement_mean.GetAchievementMeanNumber(),
// 		}

// 		// 必須パラメータチェック
// 		err := requestParamCheck(params)
// 		if err != nil {
// 			// map_str_param := interfaceToStringMap(params)
// 			// message := paramsMessageBuild(map_str_param)
// 			message := err.Error() + fmt.Sprintf(" (%+v", params) + ")"
// 			log.Println(buildInvalidParamsMessage(message))
// 			response.Response = newDefaultResponse(insufficient_param_error, message)
// 			responses.Responses = append(responses.Responses, response)
// 			return responses, nil
// 		}
// 		err = db.UpdateAchievementMean(ctx, achievement_mean)
// 		if err != nil {
// 			log.Println(err.Error())
// 			response.Response = newDefaultResponse(255, err.Error())
// 			responses.Responses = append(responses.Responses, response)
// 			return responses, nil
// 		}

// 		response.Response = newDefaultResponse(normal_code, "")
// 		responses.Responses = append(responses.Responses, response)
// 	}
// 	return responses, nil
// }

/*
/achievementMean
*/

func (*getAimoService) GetAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.GetAchievementMeanResponse, error) {
	var response *pb.GetAchievementMeanResponse = new(pb.GetAchievementMeanResponse)
	// request null validate
	// request format validate
	// get実行
	result, err := db.GetAchievementMean(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
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

func (*getAimoService) PostAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.PostAchievementMeanResponse, error) {
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

func (*getAimoService) PutAchievementMean(ctx context.Context, request *pb.AchievementMeanModel) (*pb.PutDefaultResponse, error) {
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

/*
personalEva api
*/
func (*getAimoService) GetPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (*pb.GetPersonalEvaResponse, error) {
	var response *pb.GetPersonalEvaResponse = new(pb.GetPersonalEvaResponse)

	// request null validate
	// request format validate
	// get呼び出し
	result, err := db.GetPersonalEva(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetPersonalEvaResponse_GetPersonalEvaResult)
	response.Result.PersonalEva = result

	return response, nil
}

func (*getAimoService) PostPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (*pb.PostDefaultResponse, error) {
	var response *pb.PostDefaultResponse = new(pb.PostDefaultResponse)
	// request null validate
	// request format validate
	// post呼び出し
	result, err := db.PostPersonalEva(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrRecordExist) {
			return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.PostDefaultResponse_PostResult)
	response.Result.Id = result

	return response, err
}

func (*getAimoService) PutPersonalEva(ctx context.Context, request *pb.PersonalEvaModel) (*pb.PutDefaultResponse, error) {
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)

	// request null validate
	// request format validate
	// put実行
	err := db.PutPersonalEva(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}

/*
/evaluationBefore
*/
func (*getAimoService) GetEvaluationBefore(ctx context.Context, request *pb.EvaluationBeforeModel) (*pb.GetEvaluationBeforeResponse, error) {
	var response *pb.GetEvaluationBeforeResponse = new(pb.GetEvaluationBeforeResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetEvaluationBefore(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetEvaluationBeforeResponse_GetEvaluationBeforeResult)
	response.Result.EvaluationBefore = result

	return response, nil
}

func (*getAimoService) PostEvaluationBefore(ctx context.Context, request *pb.EvaluationBeforeModel) (*pb.PostDefaultResponse, error) {
	// request null valid
	// request format valid
	// post実行
	result, err := db.PostEvaluationBefore(ctx, request)
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

func (*getAimoService) PutEvaluationBefore(ctx context.Context, request *pb.EvaluationBeforeModel) (*pb.PutDefaultResponse, error) {
	// request null valid
	// request format valid
	// put実行
	err := db.PutEvaluationBefore(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}

/*
/evaluation
*/
func (*getAimoService) GetEvaluation(ctx context.Context, request *pb.EvaluationModel) (*pb.GetEvaluationResponse, error) {
	var response *pb.GetEvaluationResponse = new(pb.GetEvaluationResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetEvaluation(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetEvaluationResponse_GetEvaluationResult)
	response.Result.Evaluation = result

	return response, nil
}

func (*getAimoService) PostEvaluation(ctx context.Context, request *pb.EvaluationModel) (*pb.PostDefaultResponse, error) {
	// request null valid
	// request format valid
	// post実行
	result, err := db.PostEvaluation(ctx, request)
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

func (*getAimoService) PutEvaluation(ctx context.Context, request *pb.EvaluationModel) (*pb.PutDefaultResponse, error) {
	// request null valid
	// request format valid
	// put実行
	err := db.PutEvaluation(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}

/*
/comprehensiveComent
*/
func (*getAimoService) GetComprehensiveComment(ctx context.Context, request *pb.ComprehensiveCommentModel) (*pb.GetComprehensiveCommentResponse, error) {
	var response *pb.GetComprehensiveCommentResponse = new(pb.GetComprehensiveCommentResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetComprehensiveComment(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetComprehensiveCommentResponse_GetComprehenesiveCommentResult)
	response.Result.ComprehensiveComment = result

	return response, nil
}

func (*getAimoService) PostComprehensiveComment(ctx context.Context, request *pb.ComprehensiveCommentModel) (*pb.PostDefaultResponse, error) {
	// request null validate
	// request format valid
	// post実行
	result, err := db.PostComprehensiveComment(ctx, request)
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

func (*getAimoService) PutComprehensiveComment(ctx context.Context, request *pb.ComprehensiveCommentModel) (*pb.PutDefaultResponse, error) {
	// request null valid
	// request format valid
	// put実行
	err := db.PutComprehensiveComment(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// response組み立て
	var response *pb.PutDefaultResponse = new(pb.PutDefaultResponse)
	response.Response = newDefaultResponse(normal_code, "")

	return response, nil
}

/*
/user
*/
func (*getAimoService) GetUser(ctx context.Context, request *pb.UserModel) (*pb.GetUserResponse, error) {
	var response *pb.GetUserResponse = new(pb.GetUserResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetUser(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
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

func (*getAimoService) PostUser(ctx context.Context, request *pb.UserModel) (*pb.PostDefaultResponse, error) {
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

func (*getAimoService) PutUser(ctx context.Context, request *pb.UserModel) (*pb.PutDefaultResponse, error) {
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

/*
/policy
*/
func (*getAimoService) GetPolicy(ctx context.Context, request *pb.PolicyModel) (*pb.GetPolicyResponse, error) {
	var response *pb.GetPolicyResponse = new(pb.GetPolicyResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetPolicy(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
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

func (*getAimoService) PostPolicy(ctx context.Context, request *pb.PolicyModel) (*pb.PostDefaultResponse, error) {
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

func (*getAimoService) PutPolicy(ctx context.Context, request *pb.PolicyModel) (*pb.PutDefaultResponse, error) {
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

/*
/departmentGoal
*/
func (*getAimoService) GetDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (*pb.GetDepartmentGoalResponse, error) {
	var response *pb.GetDepartmentGoalResponse = new(pb.GetDepartmentGoalResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetDepartmentGoal(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
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

func (*getAimoService) PostDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (*pb.PostDefaultResponse, error) {
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

func (*getAimoService) PutDepartmentGoal(ctx context.Context, request *pb.DepartmentGoalModel) (*pb.PutDefaultResponse, error) {
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

/*
/role
*/
func (*getAimoService) GetRole(ctx context.Context, request *pb.RoleModel) (*pb.GetRoleResponse, error) {
	var response *pb.GetRoleResponse = new(pb.GetRoleResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetRole(ctx, request)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
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

func (*getAimoService) PostRole(ctx context.Context, request *pb.RoleModel) (*pb.PostDefaultResponse, error) {
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

func (*getAimoService) PutRole(ctx context.Context, request *pb.RoleModel) (*pb.PutDefaultResponse, error) {
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

/*
/periods
*/
func (*getAimoService) GetPeriods(ctx context.Context, request *pb.PeriodModel) (*pb.GetPeriodsResponse, error) {
	var response *pb.GetPeriodsResponse = new(pb.GetPeriodsResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetPeriods(ctx)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetPeriodsResponse_GetPeriodsResult)
	response.Result.Periods = result

	return response, nil
}

/*
/departments
*/
func (*getAimoService) GetDepartments(ctx context.Context, request *pb.DepartmentModel) (*pb.GetDepartmentsResponse, error) {
	var response *pb.GetDepartmentsResponse = new(pb.GetDepartmentsResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetDepartments(ctx)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetDepartmentsResponse_GetDepartmentsResult)
	response.Result.Departments = result

	return response, nil
}

/*
/jobs
*/
func (*getAimoService) GetJobs(ctx context.Context, request *pb.JobModel) (*pb.GetJobsResponse, error) {
	var response *pb.GetJobsResponse = new(pb.GetJobsResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetJobs(ctx)
	if err != nil {
		if errors.Is(err, errdetails.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}
		log.Println(err)
		return nil, err
	}

	// response組み立て
	response.Response = newDefaultResponse(normal_code, "")
	response.Result = new(pb.GetJobsResponse_GetJobsResult)
	response.Result.Jobs = result

	return response, nil
}
