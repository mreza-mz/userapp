package otphandler

import (
	"context"
	"shop/core/userapp/param"
	"shop/core/userapp/validator/otpvalidator"
)

type OtpService interface {
	Send(ctx context.Context, req param.SendOTPRequest) (param.SendOTPResponse, error)
	SendForChangeUsername(ctx context.Context, req param.SendOTPRequest) (param.SendOTPResponse, error)
	Verify(req param.VerifyOTPRequest) (param.VerifyOTPResponse, error)
}

type Handler struct {
	otpSvc       OtpService
	otpValidator otpvalidator.Validator
}

func New(otpSvc OtpService, otpValidator otpvalidator.Validator) Handler {
	return Handler{
		otpSvc:       otpSvc,
		otpValidator: otpValidator,
	}
}
