package otpservice

import (
	"context"
	"shop/adapter/tracing"
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
	"shop/pkg/notifpattern"
	"shop/pkg/richerror"
	"shop/pkg/translate"
	"time"
)

func (s *Service) Send(ctx context.Context, req param.SendOTPRequest) (param.SendOTPResponse, error) {
	const op = "otpservice.Send"

	ctx, span := tracing.Service(ctx, "core-api", op)
	defer span.End()

	usernameType := entity.TypeOfUsername(req.Username)

	ttlDuration := s.config.PersistExpirationTime * time.Minute
	expDuration := s.config.ExpirationTime * time.Minute

	// we assume the repo will remove expired otp codes
	getOtp, isExist, err := s.repo.GetOTP(req.Username, usernameType)
	if err != nil {
		tracing.SpanError(span, err)
		return param.SendOTPResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if isExist {
		remainingTime := getOtp.Exp.Sub(time.Now())
		if remainingTime > (ttlDuration - expDuration) {
			return param.SendOTPResponse{ExpirationInSeconds: int((remainingTime - (ttlDuration - expDuration)).Seconds())}, nil
		}
	}

	newCode := GenerateRandomCode()
	convertedUsername := req.Username
	if usernameType == entity.PhoneNumberUsernameType {
		convertedUsername = translate.DigitToEnglish(req.Username)
	}

	persistDuration := time.Now().Add(ttlDuration)

	otp := entity.OTP{
		Username:     convertedUsername,
		UsernameType: usernameType,
		Code:         newCode,
		Exp:          persistDuration,
	}

	err = s.repo.CreateOTP(otp)
	if err != nil {
		tracing.SpanError(span, err)
		return param.SendOTPResponse{}, err
	}

	span.AddEvent("notification send started")
	if err := s.notifier(usernameType).Send(ctx, convertedUsername, "", notifpattern.SendVerifyCodePattern, []string{newCode}); err != nil {
		tracing.SpanError(span, err)
		return param.SendOTPResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	span.AddEvent("notification send finished")

	return param.SendOTPResponse{ExpirationInSeconds: int(expDuration.Seconds())}, nil
}
