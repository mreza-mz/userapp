package userservice

import (
	"context"
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
)

func (s *Service) RegisterWithOTP(ctx context.Context, req param.RegisterWithOTPReq) (param.RegisterWithOTPRes, error) {
	const op = "userservice.Register"

	usernameType := entity.TypeOfUsername(req.Username)
	user := entity.User{
		Role: entity.ManagerRole,
	}
	switch usernameType {
	case entity.PhoneNumberUsernameType:
		user.PhoneNumber = req.Username
	case entity.EmailUsernameType:
		user.Email = req.Username
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterWithOTPRes{}, err
	}

	tokens, err := s.auth.GetTokens(createdUser)
	if err != nil {
		return param.RegisterWithOTPRes{}, err
	}

	resp := param.RegisterWithOTPRes{
		User: param.UserInfo{
			ID:          createdUser.ID,
			Email:       createdUser.Email,
			PhoneNumber: createdUser.PhoneNumber,
			IsActive:    createdUser.IsActive,
			RoleID:      uint(createdUser.Role),
			Fullname:    createdUser.Fullname,
			CreatedAt:   createdUser.CreatedAt,
		},
		Tokens: tokens,
	}

	return resp, nil
}
