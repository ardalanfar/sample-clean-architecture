package route

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http/admin_http"
	"Farashop/internal/delivery/http/member_http"
	"Farashop/internal/delivery/http/public_http"
	"Farashop/internal/middlewares"
	"Farashop/pkg/validator"

	"github.com/labstack/echo/v4"
)

func Route(ech *echo.Echo, conn store.DbConn) {

	middlewares.SetMainMiddleware(ech)
	/*--------------------------------------------------------------*/
	//Public

	ech.POST("/register", public_http.Register(conn, validator.ValidateRegister()))
	ech.POST("/login", public_http.Login(conn, validator.ValidateLogin(conn)))
	ech.PATCH("/validation", public_http.MemberValidation(conn, validator.ValidateMemberValidation()))

	/*--------------------------------------------------------------*/
	//Admin Group
	AdminGroup := ech.Group("/admin")
	middlewares.SetAdminGroup(AdminGroup)

	/*------------Member Management------------*/

	MemberManagement := AdminGroup.Group("/members")
	MemberManagement.GET("/showmembers", admin_http.ShowMembers(conn))
	MemberManagement.DELETE("/deletemember/:id", admin_http.DeleteMember(conn, validator.ValidateDeleteMember(conn)))
	MemberManagement.POST("/showinfo/:id", admin_http.ShowInfoMember(conn, validator.ValidateShowInfoMember(conn)))
	/*--------------------------------------------------------------*/
	//Member Group
	MemberGroup := ech.Group("/member")
	middlewares.SetMemberGroup(MemberGroup)

	/*--------------Order Management-------------*/

	OrderManagement := MemberGroup.Group("/orders")
	OrderManagement.GET("/showorders", member_http.ShowOrders(conn))
	/*--------------------------------------------------------------*/
}
