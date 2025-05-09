package authservice

import (
	"github.com/golang-jwt/jwt/v4"
	"shop/core/userapp/entity"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID   uint        `json:"user_id"`
	Role     entity.Role `json:"role"`
	DeviceID string      `json:"device_id"`
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
