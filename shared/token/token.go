package token

import (
	"dating-apps/configs"
	"dating-apps/shared/constant"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// JWTSigningMethod is JWT's signing method
var jwtSigningMethod = jwt.SigningMethodHS256

// GenerateToken will generate both access and refresh token
// for current user.
// Access Token will be expired in 15 Minutes
// Refresh Token will be expired in 6 Months
func GenerateToken(user *UserData, params *GenerateTokenParams) (token Token, err error) {
	jwtToken, err := GenerateJWT(user, params.AccessTokenSecret, params.AccessTokenExpiry)
	if err != nil {
		return
	}

	refreshToken, e := generateRefresh(user, params.RefreshTokenSecret, params.RefreshTokenExpiry)
	if e != nil {
		return
	}

	token = Token{
		Token:        jwtToken,
		RefreshToken: refreshToken,
	}

	return
}

// GenerateJWT is
func GenerateJWT(user *UserData, tokenSecret string, tokenExpiry time.Duration) (signedToken string, err error) {
	exp := time.Now().UTC().Add(tokenExpiry)
	claims := JWTToken{
		StandardClaims: jwt.StandardClaims{
			Issuer:    configs.Get().App.Name,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   user.ID,
		},
		Username: user.Username,
		Email:    user.Email,
		OwnerID:  user.ID,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(tokenSecret))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

func generateRefresh(user *UserData, refreshTokenSecret string, refreshTokenExpiry time.Duration) (signedToken string, err error) {
	exp := time.Now().UTC().Add(refreshTokenExpiry)
	claims := JWTRefreshToken{
		StandardClaims: jwt.StandardClaims{
			Issuer:    configs.Get().App.Name,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

// ExtractClaims is
func ExtractClaims(secret, tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	log.Err(err).Msg("Invalid JWT Token")
	return nil, false
}

// GenerateVerifyEmailToken will generate reset password token
// Token will be expired in 15 Minutes
func GenerateVerifyEmailToken(userID uuid.UUID, email, username, secret string, expiry time.Duration) (signedToken string, err error) {
	exp := time.Now().UTC().Add(expiry)
	claims := JWTVerifyEmail{
		StandardClaims: jwt.StandardClaims{
			Issuer:    configs.Get().App.Name,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   userID.String(),
		},
		Email:    email,
		Username: username,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

// GetAuthenticatedUser is
func GetUserUUID(token *jwt.Token) uuid.UUID {
	claims := token.Claims.(jwt.MapClaims)

	return uuid.MustParse(claims["sub"].(string))
}

// GenerateJWTGateway is
func GenerateJWTGateway(user *UserData, tokenSecret string, tokenExpiry time.Duration) (signedToken string, err error) {
	exp := time.Now().UTC().Add(tokenExpiry)
	claims := JWTTokenGateway{
		StandardClaims: jwt.StandardClaims{
			Issuer:    configs.Get().App.Name,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   user.ID,
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(tokenSecret))
	if err != nil {
		return signedToken, err
	}

	return
}

func VerifyJwtToken(token, tokenSecret string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Err(constant.ErrInvalidAuthorization).Msg("VerifyJwtToken")
			return nil, constant.ErrInvalidAuthorization
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		log.Err(constant.ErrInvalidAuthorization).Msg("VerifyJwtToken")
		return nil, err
	}
	return jwtToken, nil
}

func ExtractTokenMetadata(token *jwt.Token) (*JWTToken, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		log.Err(constant.ErrInvalidAuthorization).Msg("ExtractTokenMetadata")
		return nil, constant.ErrTokenNotFound
	}

	name, ok := claims["name"].(string)
	if !ok {
		log.Err(constant.ErrTokenNotFound).Msg("ErrTokenNotFound")
		return nil, constant.ErrTokenNotFound
	}

	return &JWTToken{
		Username: name,
	}, nil
}
