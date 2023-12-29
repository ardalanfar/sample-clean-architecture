package validator

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateDeleteMember(store store.DbConn) contract.ValidateDeleteMember {
	return func(ctx context.Context, req dto.DeleteMemberRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.By(DoesIDExist(ctx, store))),
		)
	}
}

func ValidateShowInfoMember(store store.DbConn) contract.ValidateShowInfoMember {
	return func(ctx context.Context, req dto.ShowInfoMemberRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.By(DoesIDExist(ctx, store))),
		)
	}
}
