package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop/core/userapp/param"
	"shop/pkg/httpmsg"
)

func (h Handler) loginWithOTP(c echo.Context) error {

	var req param.LoginWithOTPReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.svc.LoginWithOTP(c.Request().Context(), req)
	if err != nil {
		return httpmsg.HttpError(err)
	}

	return c.JSON(http.StatusOK, resp)
}
