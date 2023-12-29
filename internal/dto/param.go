package dto

import "time"

/*-----------------------------------------------------*/
//Showorders model database

type Showorders struct {
	ID       uint      `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	Name     string    `json:"name" gorm:"NOT NULL,UNIQUE"`
	Number   uint      `json:"number" gorm:"NOT NULL"`
	Status   string    `json:"status" gorm:"NOT NULL"`
	Buy_time time.Time `json:"buy_time" gorm:"NOT NULL"`
}

/*-----------------------------------------------------*/
//ShowInfoMember model database

type ShowInfoMember struct {
	Username    string `json:"Username" gorm:"NOT NULL,UNIQUE"`
	Email       string `json:"Email" gorm:"NOT NULL,UNIQUE"`
	Access      uint   `json:"Access" gorm:"default:2"`
	Is_verified string `json:"Is_verified" gorm:"default:inactive" ,gorm:"NOT NULL"`
	Balance     uint   `json:"balance" gorm:"NOT NULL"`
}
