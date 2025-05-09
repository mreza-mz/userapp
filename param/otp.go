package param

type SendOTPRequest struct {
	Username string `json:"username"`
}

type SendOTPResponse struct {
	ExpirationInSeconds int `json:"interval_in_seconds"`
}
