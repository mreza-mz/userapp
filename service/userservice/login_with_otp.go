package userservice

import (
	"context"
	"shop/core/userapp/param"
)

func (s *Service) LoginWithOTP(ctx context.Context, req param.LoginWithOTPReq) (param.LoginWithOTPRes, error) {
	if _, err := s.otpVerify(param.VerifyOTPRequest{
		Username: req.Username,
		Code:     req.OTP,
	}); err != nil {
		return param.LoginWithOTPRes{}, err
	}

	user, exists, err := s.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return param.LoginWithOTPRes{}, err
	}
	if exists {
		tokens, err := s.auth.GetTokens(user)
		if err != nil {
			return param.LoginWithOTPRes{}, err
		}
		return param.LoginWithOTPRes{
			Tokens: tokens,
			UserInfo: param.UserInfo{
				Fullname:    user.Fullname,
				PhoneNumber: user.PhoneNumber,
				Email:       user.Email,
			},
			IsRegistered: false,
		}, nil
	}

	resp, rErr := s.RegisterWithOTP(ctx, param.RegisterWithOTPReq{
		Username: req.Username,
		OTP:      req.OTP,
	})

	if rErr != nil {
		return param.LoginWithOTPRes{}, rErr
	}

	return param.LoginWithOTPRes{
		Tokens:       resp.Tokens,
		UserInfo:     resp.User,
		IsRegistered: true,
	}, nil
}
