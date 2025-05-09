package otphandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/v1/otp")

	userGroup.POST("/send", h.otpSend)
	userGroup.POST("/send-username", h.otpSendForChangeUsername)
}
