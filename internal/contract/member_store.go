package contract

import (
	"Farashop/internal/dto"
	"context"
)

type MemberStore interface {
	ShowOrders(ctx context.Context, userID uint) ([]dto.Showorders, error)
}
