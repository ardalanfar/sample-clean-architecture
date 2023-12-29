package public_http

import (
	"Farashop/internal/adapter/sendmsg"
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/public_service"
	"Farashop/pkg/auth"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(conn store.DbConn, validator contract.ValidateRegister) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request  dto.RegisterUserRequest
			response dto.RegisterUserResponse
			err      error
		)

		err = json.NewDecoder(ctx.Request().Body).Decode(&request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = validator(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		response, err = public_service.NewPublic(conn).Register(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if response.Result == true {
			to := []string{
				request.Email,
			}
			sendMsg := sendmsg.NewSendMassage(to, "verify register")

			msg := sendMsg.BuildMessage()
			err = sendMsg.SendEmail(ctx.Request().Context(), msg)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func Login(conn store.DbConn, validator contract.ValidateLogin) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request  dto.LoginUserRequest
			response dto.LoginUserResponse
			err      error
		)

		err = json.NewDecoder(ctx.Request().Body).Decode(&request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = validator(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		response, err = public_service.NewPublic(conn).Login(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if response.Result {
			err = auth.GenerateTokensAndSetCookies(response.User, ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		return echo.NewHTTPError(http.StatusOK, "Wellcom "+response.User.Username)
	}
}

func MemberValidation(conn store.DbConn, validator contract.ValidateMemberValidation) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request  dto.MemberValidationRequest
			response dto.MemberValidationResponse
			err      error
		)

		err = json.NewDecoder(ctx.Request().Body).Decode(&request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = validator(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		response, err = public_service.NewPublic(conn).MemberValidation(ctx.Request().Context(), request)
		if err != nil || !response.Result {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}
