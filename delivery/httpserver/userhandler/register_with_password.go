package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop/core/userapp/param"
	"shop/pkg/httpmsg"
)

func (h Handler) registerWithPassword(c echo.Context) error {

	var req param.RegisterWithPasswordReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.svc.RegisterWithPassword(c.Request().Context(), req)
	if err != nil {
		return httpmsg.HttpError(err)
	}

	return c.JSON(http.StatusOK, resp)
}
