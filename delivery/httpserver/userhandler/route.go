package userhandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/users")

	v1.POST("/verify", h.loginWithOTP)
	v1.POST("/login", h.loginWithPassword)
	v1.POST("/register", h.registerWithPassword)
}
