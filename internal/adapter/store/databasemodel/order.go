package model

import (
	"Farashop/internal/entity"
	"time"
)

//Order database model

type Order struct {
	ID         uint      `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	User_id    uint      `json:"user_id" gorm:"NOT NULL"`
	Product_id uint      `json:"product_id" gorm:"NOT NULL"`
	Number     uint      `json:"number" gorm:"NOT NULL"`
	Status     string    `json:"status" gorm:"NOT NULL"`
	Buy_time   time.Time `json:"buy_time" gorm:"NOT NULL"`
}

/*-----------------------------------------------------*/
//convert data model to database model

func MapFromOrderEntity(order entity.Order) Order {
	return Order{
		ID:       order.ID,
		Number:   order.Number,
		Buy_time: order.Buy_time,
		Status:   order.Status,
	}
}

func MapToOrderEntity(order Order) entity.Order {
	return entity.Order{
		ID:       order.ID,
		Number:   order.Number,
		Buy_time: order.Buy_time,
		Status:   order.Status,
	}
}

/*-----------------------------------------------------*/
