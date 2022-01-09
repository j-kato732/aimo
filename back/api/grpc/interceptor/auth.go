package interceptor

import (
	"context"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	log.Println(token)

	return ctx, nil
}
