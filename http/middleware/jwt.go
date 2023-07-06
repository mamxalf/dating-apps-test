package middleware

import (
	"context"
	"dating-apps/configs"
	"dating-apps/shared/failure"
	"dating-apps/shared/response"
	"dating-apps/shared/token"
	"net/http"
	"strings"
)

type JwtKeyContext string

const (
	JwtKeyContextClaims JwtKeyContext = "jwt-claims"
)

type JWT struct {
	Config *configs.Config
}

func ProvideJWTMiddleware(config *configs.Config) *JWT {
	return &JWT{
		Config: config,
	}
}

func (j *JWT) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authorization) <= 1 {
			response.WithError(w, failure.Unauthorized(failure.ErrUnauthorized.Error()))
			return
		}

		jwtToken, err := token.VerifyJwtToken(authorization[1], j.Config.Auth.AccessTokenSecret)
		if err != nil {
			response.WithError(w, failure.Unauthorized(failure.ErrUnauthorized.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), JwtKeyContextClaims, jwtToken.Claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaimsUser(ctx context.Context) interface{} {
	return ctx.Value(JwtKeyContextClaims)
}
