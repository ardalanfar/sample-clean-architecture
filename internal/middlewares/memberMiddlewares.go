package middlewares

import (
	"Farashop/pkg/auth"
	"Farashop/pkg/consts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMemberGroup(grp *echo.Group) {
	grp.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(consts.JwtSecretKey),
		TokenLookup:             "cookie:" + consts.AccessTokenCookieName, // "<source>:<name>"
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))

	grp.Use(TokenRefresherMiddleware)
	grp.Use(CheckAccessMember)
}

