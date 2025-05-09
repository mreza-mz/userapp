package param

type LoginWithOTPReq struct {
	Username string `json:"username"`
	OTP      string `json:"otp"`
}

type LoginWithOTPRes struct {
	Tokens       Tokens   `json:"tokens"`
	UserInfo     UserInfo `json:"user_info"`
	IsRegistered bool     `json:"is_registered"`
}
type LoginWithPasswordReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginWithPasswordRes struct {
	Tokens   Tokens   `json:"tokens"`
	UserInfo UserInfo `json:"user_info"`
}
