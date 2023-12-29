package contract

import (
	"Farashop/internal/entity"
	"context"
)

type PublicStore interface {
	Register(ctx context.Context, user entity.User) (bool, error)
	Login(ctx context.Context, user entity.User) (entity.User, error)
	MemberValidation(ctx context.Context, user entity.User) (bool, error)
}
