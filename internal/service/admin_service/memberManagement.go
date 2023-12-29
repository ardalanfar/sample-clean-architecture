package admin_service

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"
)

type Interactor struct {
	store contract.AdminStore
}

func NewAdmin(store contract.AdminStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) ShowMembers(ctx context.Context, _ dto.ShowMembersRequest) (dto.ShowMembersResponse, error) {
	users, err := i.store.ShowMembers(ctx)
	if err != nil {
		return dto.ShowMembersResponse{}, err
	}
	return dto.ShowMembersResponse{Users: users}, nil
}

func (i Interactor) DeleteMember(ctx context.Context, req dto.DeleteMemberRequest) (dto.DeleteMemberResponse, error) {
	err := i.store.DeleteMember(ctx, req.ID)
	if err != nil {
		return dto.DeleteMemberResponse{Result: false}, err
	}
	return dto.DeleteMemberResponse{Result: true}, nil
}

func (i Interactor) ShowInfoMember(ctx context.Context, req dto.ShowInfoMemberRequest) (dto.ShowInfoMemberResponse, error) {
	info, err := i.store.ShowInfoMember(ctx, req.ID)
	if err != nil {
		return dto.ShowInfoMemberResponse{Info: info}, err
	}
	return dto.ShowInfoMemberResponse{Info: info}, nil
}
