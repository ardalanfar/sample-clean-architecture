package middlewares

import (
	"Farashop/pkg/auth"
	"Farashop/pkg/consts"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckAccessMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("user") == nil {
			return next(ctx)
		}
		u := ctx.Get("user").(*jwt.Token)
		claims := u.Claims.(*auth.Claims)
		if int(claims.Access) == consts.MemberAccess {
			return next(ctx)
		}
		return ctx.JSON(http.StatusBadRequest, "You do not have access")
	}
}

func CheckAccessAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("user") == nil {
			return next(ctx)
		}
		u := ctx.Get("user").(*jwt.Token)
		claims := u.Claims.(*auth.Claims)
		if int(claims.Access) == consts.AdminAccess {
			return next(ctx)
		}
		return ctx.JSON(http.StatusBadRequest, "You do not have access")
	}
}
