package userservice

import (
	"context"
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
)

func (s *Service) RegisterWithPassword(ctx context.Context, req param.RegisterWithPasswordReq) (param.RegisterWithPasswordRes, error) {
	const op = "userservice.Register"

	_, exists, err := s.GetUserByUsername(ctx, req.Username)
	// error happened
	if err != nil {
		return param.RegisterWithPasswordRes{}, err
	}
	// user was already existing with this username
	if exists {
		return param.RegisterWithPasswordRes{}, richerror.New(op).WithMessage(errmsg.ErrorMsgPhoneNumberDuplicated).WithKind(richerror.KindBadRequest)
	}

	usernameType := entity.TypeOfUsername(req.Username)
	password, _ := entity.PasswordHash(req.Password)
	user := entity.User{
		Password: password,
		Fullname: req.Fullname,
		Role:     entity.TenantRole,
	}
	switch usernameType {
	case entity.PhoneNumberUsernameType:
		user.PhoneNumber = req.Username
	case entity.EmailUsernameType:
		user.Email = req.Username
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterWithPasswordRes{}, err
	}

	tokens, err := s.auth.GetTokens(createdUser)
	if err != nil {
		return param.RegisterWithPasswordRes{}, err
	}

	resp := param.RegisterWithPasswordRes{
		User: param.UserInfo{
			ID:                createdUser.ID,
			Email:             createdUser.Email,
			PhoneNumber:       createdUser.PhoneNumber,
			IsActive:          createdUser.IsActive,
			RoleID:            uint(createdUser.Role),
			IsChangedPassword: createdUser.IsChangedPassword,
			Fullname:          createdUser.Fullname,
			CreatedAt:         createdUser.CreatedAt,
		},
		Tokens: tokens,
	}
	return resp, nil
}
