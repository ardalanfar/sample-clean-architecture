package entity

type User struct {
	ID                uint   `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Access            uint   `json:"access"`
	Verification_code uint   `json:"verification_code"`
	Is_verified       string `json:"is_verified"`
}
