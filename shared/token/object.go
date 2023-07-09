package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GenerateTokenParams struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// JWTToken is
type JWTToken struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	OwnerID  string `json:"owner_id"`
	jwt.StandardClaims
}

// JWTTokenGateway is
type JWTTokenGateway struct {
	Msisdn         string `json:"msisdn"`
	Username       string `json:"username"`
	UserCategoryID int32  `json:"user_category_id"`
	jwt.StandardClaims
}

// JWTRefreshToken is
type JWTRefreshToken struct {
	jwt.StandardClaims
}

// JWTVerifyEmail is
type JWTVerifyEmail struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	Username string `json:"username"`
}

type InterceptorClaims struct {
	UserID         uuid.UUID
	UserCategoryID int32
}
