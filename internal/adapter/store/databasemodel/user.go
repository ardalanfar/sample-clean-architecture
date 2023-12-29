package model

import "Farashop/internal/entity"

//User database model

type User struct {
	ID                uint   `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	Username          string `json:"username" gorm:"NOT NULL,UNIQUE"`
	Email             string `json:"email" gorm:"NOT NULL,UNIQUE"`
	Password          string `json:"password" gorm:"NOT NULL"`
	Access            uint   `json:"access" gorm:"default:2"`
	Verification_code uint   `json:"verification_code" gorm:"NOT NULL,UNIQUE"`
	Is_verified       string `json:"is_verified" gorm:"default:inactive" ,gorm:"NOT NULL"`
}

/*-----------------------------------------------------*/
//convert data model to database model

func MapFromUserEntity(user entity.User) User {
	return User{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Access:            user.Access,
		Verification_code: user.Verification_code,
		Is_verified:       user.Is_verified,
	}
}

func MapToUserEntity(user User) entity.User {
	return entity.User{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Access:            user.Access,
		Verification_code: user.Verification_code,
		Is_verified:       user.Is_verified,
	}
}

/*-----------------------------------------------------*/
