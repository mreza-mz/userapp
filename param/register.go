package param

type RegisterWithPasswordReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	DeviceID string `json:"device_id"`
}

type RegisterWithPasswordRes struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}

type RegisterWithOTPReq struct {
	Username string `json:"username"`
	OTP      string `json:"otp"`
}
type RegisterWithOTPRes struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}
