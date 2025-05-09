package userservice

import (
	"context"
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
)

func (s *Service) LoginWithPassword(ctx context.Context, req param.LoginWithPasswordReq) (param.LoginWithPasswordRes, error) {
	const op = "userservice.LoginWithPassword"

	user, exists, err := s.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return param.LoginWithPasswordRes{}, err
	}
	if !exists {
		return param.LoginWithPasswordRes{}, richerror.New(op).WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindBadRequest)
	}

	err = entity.CheckPasswordHash(user.Password, req.Password)
	if err != nil {
		return param.LoginWithPasswordRes{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgPasswordIsNotValid).WithKind(richerror.KindInvalid)
	}

	tokens, err := s.auth.GetTokens(user)
	if err != nil {
		return param.LoginWithPasswordRes{}, err
	}

	return param.LoginWithPasswordRes{
		UserInfo: param.UserInfo{
			ID:                user.ID,
			Email:             user.Email,
			PhoneNumber:       user.PhoneNumber,
			Fullname:          user.Fullname,
			IsChangedPassword: user.IsChangedPassword,
			RoleID:            uint(user.Role),
		},
		Tokens: tokens,
	}, nil
}
