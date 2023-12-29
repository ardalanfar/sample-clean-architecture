package contract

import (
	"Farashop/internal/dto"
	"context"
)

type MemberInteractor interface {
	ShowOrders(context.Context, dto.ShowOrdersRequest) (dto.ShowOrdersResponse, error)
}
