package mockotpservice

import (
	"context"
	"fmt"
	"shop/core/userapp/param"
	"shop/pkg/errmsg"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s Service) Send(ctx context.Context, req param.SendOTPRequest) (param.SendOTPResponse, error) {
	return param.SendOTPResponse{
		ExpirationInSeconds: 120,
	}, nil
}

func (s Service) SendForChangeUsername(ctx context.Context, req param.SendOTPRequest) (param.SendOTPResponse, error) {
	return param.SendOTPResponse{
		ExpirationInSeconds: 120,
	}, nil
}

func (s Service) Verify(req param.VerifyOTPRequest) (param.VerifyOTPResponse, error) {
	const constantOTPCode = "202020"
	if constantOTPCode != req.Code {
		return param.VerifyOTPResponse{}, fmt.Errorf(errmsg.ErrorMsgOTPIsNotValid)
	}
	return param.VerifyOTPResponse{}, nil
}
