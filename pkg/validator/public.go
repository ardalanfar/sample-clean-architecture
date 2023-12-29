package validator

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateRegister() contract.ValidateRegister {
	return func(ctx context.Context, req dto.RegisterUserRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.Username, validation.Required),
			validation.Field(&req.Email, validation.Required, is.Email),
			validation.Field(&req.Password, validation.Required),
		)
	}
}

func ValidateLogin(conn store.DbConn) contract.ValidateLogin {
	return func(ctx context.Context, req dto.LoginUserRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.Username, validation.By(DoesUsernameActive(ctx, conn))),
			validation.Field(&req.Password, validation.Required),
		)
	}
}

func ValidateMemberValidation() contract.ValidateMemberValidation {
	return func(ctx context.Context, req dto.MemberValidationRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.Username, validation.Required),
			validation.Field(&req.Code, validation.Required),
		)
	}
}
