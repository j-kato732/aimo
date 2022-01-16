package interceptor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomClaims struct {
	Audience         []string `json:"aud"`
	AuthorizedPartys string   `json:"azp"`
	Scope            string   `json:"scope"`
	Permissions      []string `json:"permissions"`
	jwt.StandardClaims
}

type (
	jwks struct {
		Keys []jsonWebKeys `json:"keys"`
	}

	jsonWebKeys struct {
		Kty string   `json:"kty"`
		Kid string   `json:"kid"`
		Use string   `json:"use"`
		N   string   `json:"n"`
		E   string   `json:"e"`
		X5c []string `json:"x5c"`
	}
)

func loadEnv() error {

	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	//もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		return err
	}

	return nil
}

func AuthFuncHandler(ctx context.Context) (context.Context, error) {
	ctx, err := authFunc(ctx)
	if err != nil {
		log.Println(fmt.Errorf("Authentication Error: %w", err))
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}

	return ctx, nil
}

func authFunc(ctx context.Context) (context.Context, error) {
	err := loadEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to loadEnv: %w", err)
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, fmt.Errorf("failed to AuthFromMD: %w", err)
	}
	log.Println(token)

	keyfunc := func(token *jwt.Token) (interface{}, error) {
		certificate, err := getPEMCertificate(token)
		if err != nil {
			return nil, fmt.Errorf("failed to getPEMCertficate: %w", err)
		}

		return jwt.ParseRSAPublicKeyFromPEM([]byte(certificate))
	}

	tokenWithClaim, err := jwt.ParseWithClaims(token, &CustomClaims{}, keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to ParseWithClaims: %w", err)
	}

	if tokenWithClaim.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	err = tokenWithClaim.Claims.Valid()
	if err != nil {
		return nil, fmt.Errorf("failed to tokenWithClaims.Claims.Valid(): %w", err)
	}

	return ctx, nil
}

func (c CustomClaims) Valid() error {
	expectedAudience := os.Getenv("AUTH0_AUDIENCE")
	if c.Audience[0] != expectedAudience {
		return fmt.Errorf("token claims validation failed: unexpected audience %q", c.Audience)
	}

	expectedIssuer := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	if c.Issuer != expectedIssuer {
		return fmt.Errorf("token claims validation failed: unexpected issuer %q", c.Issuer)
	}

	err := c.StandardClaims.Valid()
	if err != nil {
		return fmt.Errorf("failed to StandardClaims.Valid(): %w", err)
	}

	return nil
}

func getPEMCertificate(token *jwt.Token) (string, error) {
	response, err := http.Get("https://" + os.Getenv("AUTH0_DOMAIN") + "/.well-known/jwks.json")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var jwks jwks
	if err = json.NewDecoder(response.Body).Decode(&jwks); err != nil {
		return "", err
	}

	var cert string
	for _, key := range jwks.Keys {
		if token.Header["kid"] == key.Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"
			break
		}
	}

	if cert == "" {
		return cert, errors.New("unable to find appropriate key")
	}

	return cert, nil
}
