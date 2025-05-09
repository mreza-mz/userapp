package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID                uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
	LastLogin         *time.Time
	PhoneNumber       string
	Email             string
	Fullname          string
	Password          string
	Avatar            string
	Role              Role
	IsChangedPassword bool
	IsActive          bool
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(currentPassword string, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(requestPassword))
	return err
}
