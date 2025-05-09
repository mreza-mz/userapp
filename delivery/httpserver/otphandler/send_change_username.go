package otphandler

import (
	"net/http"
	"shop/core/userapp/param"
	"shop/pkg/errmsg"
	"shop/pkg/httpmsg"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h Handler) otpSendForChangeUsername(c echo.Context) error {
	var req param.SendOTPRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	//if fieldErrors, err := h.otpValidator.ValidateSendRequest(c.Request().Context(), req); err != nil {
	//	msg, code := httpmsg.Error(err)
	//	return c.JSON(code, echo.Map{
	//		"message": errmsg.TranslateFa(msg),
	//		"errors":  fieldErrors,
	//	})
	//}

	req.Username = strings.Trim(req.Username, " ")
	req.Username = strings.ToLower(req.Username)

	res, err := h.otpSvc.SendForChangeUsername(c.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, errmsg.TranslateFa(msg))
	}

	return c.JSON(http.StatusOK, res)
}
