package otpservice

import (
	"fmt"
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
	"shop/pkg/translate"
)

func (s *Service) Verify(req param.VerifyOTPRequest) (param.VerifyOTPResponse, error) {
	// get usernameType
	usernameType := entity.TypeOfUsername(req.Username)

	// if type is phone ? convert code to english number
	convertedUsername := req.Username
	if usernameType == "phone" {
		convertedUsername = translate.DigitToEnglish(req.Username)
	}

	// get otp by username from repo
	rowOtp, isExist, err := s.repo.GetOTP(convertedUsername, usernameType)
	if err != nil {
		return param.VerifyOTPResponse{}, richerror.New("unknown").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if !isExist {
		return param.VerifyOTPResponse{}, fmt.Errorf(errmsg.ErrorMsgOTPIsNotFound)
	}

	if rowOtp.Code != req.Code {
		return param.VerifyOTPResponse{}, fmt.Errorf(errmsg.ErrorMsgOTPIsNotValid)
	}

	err = s.repo.DeleteOTP(convertedUsername, usernameType)
	if err != nil {
		return param.VerifyOTPResponse{}, richerror.New("unknown").WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return param.VerifyOTPResponse{}, nil
}
