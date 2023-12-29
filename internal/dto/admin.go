package dto

import "Farashop/internal/entity"

/*----------------Show Members--------------------*/

type ShowMembersRequest struct{}

type ShowMembersResponse struct {
	Users []entity.User `json:"users"`
}

/*---------------Delete Member---------------------*/

type DeleteMemberRequest struct {
	ID uint `json:"id"`
}

type DeleteMemberResponse struct {
	Result bool `json:"result"`
}

/*----------------Show Information---------------*/

type ShowInfoMemberRequest struct {
	ID uint `json:"id"`
}

type ShowInfoMemberResponse struct {
	Info ShowInfoMember
}
