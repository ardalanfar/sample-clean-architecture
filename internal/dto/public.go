package dto

import "Farashop/internal/entity"

/*-----------------Register user----------------------*/

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Result bool `json:"result"`
}

/*--------------------Login user----------------------*/

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Result bool        `json:"result"`
	User   entity.User `json:"user"`
}

/*---------------------Member Validation----------------*/

type MemberValidationRequest struct {
	Username string `json:"username"`
	Code     uint   `json:"code"`
}

type MemberValidationResponse struct {
	Result bool `json:"result"`
}
