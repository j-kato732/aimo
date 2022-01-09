package main

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/j-kato732/aimo/proto"
)

const (
	normal_code = 1

	insufficient_param_error = 10
	invalid_param_format     = 11
)

var (
	response_status  int64
	response_message string
)

func newDefaultResponse(status int64, message string) *pb.DefaultResponse {
	default_response := new(pb.DefaultResponse)
	default_response.Status = status
	default_response.Message = message
	return default_response
}

func buildInvalidParamsMessage(message string) string {
	return "invalid require params " + message
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
