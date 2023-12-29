package store

import (
	model "Farashop/internal/adapter/store/databasemodel"
	"Farashop/pkg/encrypt"

	"gorm.io/gorm"
)

func InsertSeedAdmin(Db *gorm.DB) error {
	var (
		err  error
		user model.User
	)

	res := Db.Where("username = ?", "admin").Find(&user)
	if res.RowsAffected != 0 {
		return nil
	}

	pass, _ := encrypt.HashPassword("123456")
	user = model.User{Username: "admin", Email: "admin@yahoo.com", Password: pass, Access: 1, Is_verified: "active"}

	err = Db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
