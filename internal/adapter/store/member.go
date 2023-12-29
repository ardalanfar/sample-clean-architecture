package store

import (
	"Farashop/internal/dto"
	"context"
)

func (conn DbConn) ShowOrders(ctx context.Context, userID uint) ([]dto.Showorders, error) {
	var (
		Showorders []dto.Showorders
	)

	cheek := conn.Db.WithContext(ctx).Table("orders").Select("orders.id", "orders.number", "orders.buy_time", "orders.status", "products.name").Joins("JOIN products ON products.id = orders.product_id").Where("user_id", userID).Find(&Showorders)
	if cheek.Error != nil {
		return nil, cheek.Error
	}

	return Showorders, nil
}
