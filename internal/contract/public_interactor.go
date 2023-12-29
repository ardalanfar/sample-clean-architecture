package contract

import (
	"Farashop/internal/dto"
	"context"
)

type PublicInteractor interface {
	Register(context.Context, dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
	Login(context.Context, dto.LoginUserRequest) (dto.LoginUserResponse, error)
	MemberValidation(context.Context, dto.MemberValidationRequest) (dto.MemberValidationResponse, error)
}
