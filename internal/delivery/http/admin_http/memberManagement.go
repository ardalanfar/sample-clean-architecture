package admin_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/admin_service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ShowMembers(conn store.DbConn) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request  dto.ShowMembersRequest
			response dto.ShowMembersResponse
			err      error
		)

		response, err = admin_service.NewAdmin(conn).ShowMembers(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, response)
	}
}

func DeleteMember(conn store.DbConn, validator contract.ValidateDeleteMember) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request  dto.DeleteMemberRequest
			response dto.DeleteMemberResponse
			err      error
		)

		userID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		request = dto.DeleteMemberRequest{ID: uint(userID)}

		err = validator(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		response, err = admin_service.NewAdmin(conn).DeleteMember(ctx.Request().Context(), request)
		if err != nil && !response.Result {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func ShowInfoMember(conn store.DbConn, validator contract.ValidateShowInfoMember) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			err      error
			request  dto.ShowInfoMemberRequest
			response dto.ShowInfoMemberResponse
		)

		userID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		request = dto.ShowInfoMemberRequest{ID: uint(userID)}

		err = validator(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		response, err = admin_service.NewAdmin(conn).ShowInfoMember(ctx.Request().Context(), request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, response)
	}
}
