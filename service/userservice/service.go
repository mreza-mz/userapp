package userservice

import (
	"context"
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
	"shop/core/userapp/service/authservice"
)

type Repository interface {
	GetUserExistByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, bool, error)
	GetUserExistByEmail(ctx context.Context, email string) (entity.User, bool, error)
	Register(u entity.User) (entity.User, error)
}

type OTPVerifier func(req param.VerifyOTPRequest) (param.VerifyOTPResponse, error)

type Service struct {
	auth      authservice.Service
	repo      Repository
	otpVerify OTPVerifier
}

func New(auth authservice.Service, repository Repository) Service {
	return Service{auth: auth, repo: repository}
}

func (s *Service) WithOtpService(otpSvc OTPVerifier) {
	s.otpVerify = otpSvc
}
