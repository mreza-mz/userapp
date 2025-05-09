package authservice

import (
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	SignKey               string        `koanf:"sign_key"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string        `koanf:"refresh_subject"`
}

type AuthenticatedHandler struct {
	config  Config
	service Service
}

func NewAuthenticatedHandler(config Config, service Service) AuthenticatedHandler {
	return AuthenticatedHandler{
		config:  config,
		service: service,
	}
}

func (h AuthenticatedHandler) SignKey() []byte {
	return []byte(h.config.SignKey)
}

func (h AuthenticatedHandler) ParseToken(bearerToken string) (interface{}, error) {
	return h.service.ParseToken(bearerToken)
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

func (s Service) GetTokens(user entity.User) (param.Tokens, error) {
	accessToken, err := s.createAccessToken(user)
	if err != nil {
		return param.Tokens{}, err
	}
	refreshToken, err := s.createRefreshToken(user)
	if err != nil {
		return param.Tokens{}, err
	}
	return param.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s Service) createAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) createRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (s Service) createToken(userID uint, role entity.Role, subject string, expireDuration time.Duration) (string, error) {
	// create a signer for rsa 256
	// TODO: replace with rsa 256 RS256 - https://github.com/golang-jwt/jwt/blob/main/http_example_test.go

	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
		Role:   role,
	}

	// TODO: add sign method to config
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
