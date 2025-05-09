package entity

import (
	"regexp"
	"time"
)

type UsernameType string

const (
	EmailUsernameType       UsernameType = "email"
	PhoneNumberUsernameType UsernameType = "phone"
)

type OTP struct {
	Username     string       `json:"username"`
	UsernameType UsernameType `json:"username_type"`
	Code         string       `json:"code"`
	Exp          time.Time    `json:"exp"`
}

func TypeOfUsername(username string) UsernameType {
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneNumberPattern := regexp.MustCompile(`^09[0-9-]+$`)

	if emailPattern.MatchString(username) {
		return EmailUsernameType
	}

	if phoneNumberPattern.MatchString(username) {
		return PhoneNumberUsernameType
	}

	return ""
}
