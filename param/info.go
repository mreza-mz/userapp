package param

import "time"

type UserInfo struct {
	ID                uint      `json:"id"`
	PhoneNumber       string    `json:"phone_number"`
	Email             string    `json:"email"`
	RoleID            uint      `json:"role_id"`
	Fullname          string    `json:"fullname"`
	IsChangedPassword bool      `json:"is_changed_password"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
}
