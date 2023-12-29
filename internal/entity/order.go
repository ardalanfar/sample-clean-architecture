package entity

import "time"

type Order struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Number   uint      `json:"number"`
	Buy_time time.Time `json:"buy_time"`
	Status   string    `json:"status"`
}
