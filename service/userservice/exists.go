package userservice

import (
	"context"
	"shop/core/userapp/entity"
	"shop/pkg/richerror"
)

func (s *Service) GetUserByUsername(ctx context.Context, username string) (entity.User, bool, error) {
	const op = "userservice.GetUserExist"

	usernameType := entity.TypeOfUsername(username)

	switch usernameType {
	case entity.PhoneNumberUsernameType:
		user, exists, eErr := s.repo.GetUserExistByPhoneNumber(ctx, username)
		if eErr != nil {
			return entity.User{}, false, eErr
		}
		return user, exists, nil

	case entity.EmailUsernameType:
		user, exists, eErr := s.repo.GetUserExistByEmail(ctx, username)
		if eErr != nil {
			return entity.User{}, false, eErr
		}
		return user, exists, nil
	default:
		return entity.User{}, false, richerror.New(op).WithMessage("invalid username type").
			WithMeta(map[string]interface{}{"username": username})
	}
}
