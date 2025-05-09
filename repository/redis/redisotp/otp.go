package redisotp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"shop/core/userapp/entity"
	"shop/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

const otpKeyPrefix = "otp"

func otpKey(usernameType entity.UsernameType, username string) string {
	return fmt.Sprintf("%s:%s:%s", otpKeyPrefix, usernameType, username)
}

func (d DB) CreateOTP(otp entity.OTP) error {
	// TODO: pass context from handler
	err := d.adapter.Client().Set(context.Background(), otpKey(otp.UsernameType, otp.Username), otp.Code, otp.Exp.Sub(time.Now())).Err()
	if err != nil {
		return fmt.Errorf("error adding OTP to redis: %w", err)
	}

	return nil
}

func (d DB) DeleteOTP(username string, usernameType entity.UsernameType) error {
	err := d.adapter.Client().Del(context.Background(), otpKey(usernameType, username)).Err()
	if err != nil {
		return fmt.Errorf("error deleting OTP from redis: %w", err)
	}
	return nil
}

func (d DB) GetOTP(username string, usernameType entity.UsernameType) (entity.OTP, bool, error) {
	// TODO: pass context from handler
	value, err := d.adapter.Client().Get(context.Background(), otpKey(usernameType, username)).Result()
	if errors.Is(err, redis.Nil) {
		return entity.OTP{}, false, nil
	}

	logger.L().Debug("redisotp.GetOTP", slog.Any("error", err), slog.String("key", otpKey(usernameType, username)))

	if err != nil {
		return entity.OTP{}, false, fmt.Errorf("error get OTP from redis: %w", err)
	}

	ttlDuration, err := d.adapter.Client().TTL(context.Background(), otpKey(usernameType, username)).Result()
	if err != nil {
		return entity.OTP{}, false, fmt.Errorf("error getting TTL from redis: %w", err)
	}

	//remainingSeconds := int64(ttlDuration.Seconds())

	otp := entity.OTP{
		UsernameType: usernameType,
		Username:     username,
		Code:         value,
		Exp:          time.Now().Add(ttlDuration),
	}

	return otp, true, nil
}
