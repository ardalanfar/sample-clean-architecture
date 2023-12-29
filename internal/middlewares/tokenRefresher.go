package middlewares

import (
	"Farashop/internal/entity"
	"Farashop/pkg/auth"
	"Farashop/pkg/consts"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)


func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("user") == nil {
			return next(ctx)
		}
		u := ctx.Get("user").(*jwt.Token)
		claims := u.Claims.(*auth.Claims)

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < (10 * time.Minute) {
			rc, err := ctx.Cookie(consts.RefreshTokenCookieName)
			if err == nil && rc != nil {
				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(consts.JwtRefreshSecretKey), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						ctx.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					_ = auth.GenerateTokensAndSetCookies(entity.User{
						Username: claims.Name,
						ID:       claims.ID,
						Access:   claims.Access,
					}, ctx)
				}
			}
		}
		return next(ctx)
	}
}