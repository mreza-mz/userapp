package otpservice

import (
	"math/rand"
	"shop/core/userapp/entity"
	"shop/core/userapp/service/userservice"
	"shop/pkg/notifier"
	"strconv"
	"time"
)

type Repository interface {
	CreateOTP(otp entity.OTP) error
	GetOTP(username string, usernameType entity.UsernameType) (entity.OTP, bool, error)
	DeleteOTP(username string, usernameType entity.UsernameType) error
}

type Config struct {
	ExpirationTime        time.Duration `koanf:"expiration_time"`
	PersistExpirationTime time.Duration `koanf:"persist_expiration_time"`
	UseMock               bool          `koanf:"use_mock"`
}

type Service struct {
	config   Config
	notifier func(usernameType entity.UsernameType) notifier.Notifier
	repo     Repository
	userSvc  userservice.Service
}

func New(cfg Config, notifier func(usernameType entity.UsernameType) notifier.Notifier, repo Repository,
	userSvc userservice.Service) Service {
	return Service{config: cfg, notifier: notifier, repo: repo, userSvc: userSvc}
}

func GenerateRandomCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
