package otpvalidator

import (
	"context"
	"shop/adapter/tracing"
	"shop/core/userapp/entity"
	"shop/core/userapp/param"
	"shop/pkg/errmsg"
	"shop/pkg/regexpatterns"
	"shop/pkg/richerror"

	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateSendRequest(ctx context.Context, req param.SendOTPRequest) (map[string]string, error) {
	const op = "otpvalidator.ValidateRegisterRequest"

	ctx, span := tracing.Validator(ctx, "core-api", op)
	defer span.End()

	usernameType := entity.TypeOfUsername(req.Username)

	switch usernameType {
	case entity.PhoneNumberUsernameType:
		if err := validation.ValidateStruct(&req,
			validation.Field(&req.Username,
				validation.Required,
				validation.Length(3, 50)),

			validation.Field(&req.Username,
				validation.Required,
				validation.Match(regexp.MustCompile(regexpatterns.PhoneNumber)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
			),
		); err != nil {
			fieldErrors := make(map[string]string)

			errV, ok := err.(validation.Errors)
			if ok {
				for key, value := range errV {
					if value != nil {
						fieldErrors[key] = value.Error()
					}
				}
			}

			return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
				WithKind(richerror.KindInvalid).
				WithMeta(map[string]interface{}{"req": req}).WithErr(err)
		}

	case entity.EmailUsernameType:
		if err := validation.ValidateStruct(&req,
			validation.Field(&req.Username,
				validation.Required,
			),

			validation.Field(&req.Username,
				validation.Required,
				validation.Match(regexp.MustCompile(regexpatterns.EmailRegex)).Error(errmsg.ErrorMsgEmailIsNotValid),
			),
		); err != nil {
			fieldErrors := make(map[string]string)

			errV, ok := err.(validation.Errors)
			if ok {
				for key, value := range errV {
					if value != nil {
						fieldErrors[key] = value.Error()
					}
				}
			}

			return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
				WithKind(richerror.KindInvalid).
				WithMeta(map[string]interface{}{"req": req}).WithErr(err)
		}
	default:
		return map[string]string{"username": errmsg.ErrorMsgInvalidInput}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req})
	}

	return nil, nil
}
