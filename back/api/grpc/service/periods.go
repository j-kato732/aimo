package service

import (
	"context"
	"errors"
	"log"

	db "github.com/j-kato732/aimo/db"
	errdetails "github.com/j-kato732/aimo/errors"
	pb "github.com/j-kato732/aimo/proto"
)

func (*AimoService) GetPeriods(ctx context.Context, request *pb.PeriodModel) (*pb.GetPeriodsResponse, error) {
	var response *pb.GetPeriodsResponse = new(pb.GetPeriodsResponse)

	// request null valid
	// request format valid
	// get実行
	result, err := db.GetPeriods(ctx)
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
	response.Result = new(pb.GetPeriodsResponse_GetPeriodsResult)
	response.Result.Periods = result

	return response, nil
}
