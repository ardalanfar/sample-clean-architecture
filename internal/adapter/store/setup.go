package store

import (
	model "Farashop/internal/adapter/store/databasemodel"
	"Farashop/internal/config"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConn struct {
	Db *gorm.DB
}

func New(cfg *config.Config) (DbConn, error) {
	var (
		err      error
		database *gorm.DB
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DB.Host,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Dbname,
	)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return DbConn{}, errors.Wrap(err, "Failed to connect to database!")
	}
	return DbConn{
		Db: database,
	}, nil
}

func (db DbConn) Migratate() error {
	var (
		err error
	)
	err = db.Db.AutoMigrate(&model.User{}, &model.Order{})
	if err != nil {
		panic("Failed to auto migrate database!")
	}

	err = InsertSeedAdmin(db.Db)
	if err != nil {
		panic("Failed to Insert Admin")
	}

	return errors.Wrap(err, "Failed to auto migrate database!")
}
