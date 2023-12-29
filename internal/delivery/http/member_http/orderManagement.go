package member_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/dto"
	"Farashop/internal/service/member_service"
	"Farashop/pkg/auth"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ShowOrders(conn store.DbConn) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			response dto.ShowOrdersResponse
			request  dto.ShowOrdersRequest
			err      error
		)

		usrCtx := ctx.Get("user").(*jwt.Token)
		claims := usrCtx.Claims.(*auth.Claims)
		request.ID = claims.ID

		response, err = member_service.NewMember(conn).ShowOrders(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
