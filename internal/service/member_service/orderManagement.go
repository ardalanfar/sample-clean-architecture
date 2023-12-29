package member_service

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"
)

type Interactor struct {
	store contract.MemberStore
}

func NewMember(store contract.MemberStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) ShowOrders(ctx context.Context, req dto.ShowOrdersRequest) (dto.ShowOrdersResponse, error) {
	info, err := i.store.ShowOrders(ctx, req.ID)
	if err != nil {
		return dto.ShowOrdersResponse{}, err
	}
	return dto.ShowOrdersResponse{Orders: info}, nil
}
