package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

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
