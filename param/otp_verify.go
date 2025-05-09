package param

type VerifyOTPRequest struct {
	Username string
	Code     string
}

type VerifyOTPResponse struct{}
