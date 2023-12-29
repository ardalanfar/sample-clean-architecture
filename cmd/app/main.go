package main

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/config"
	"Farashop/pkg/route"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	var (
		err  error
		ech  *echo.Echo
		conn store.DbConn
		cfg  *config.Config
	)

	cfg, err = config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}

	conn, err = store.New(cfg)
	if err != nil {
		fmt.Println(err)
	}

	err = conn.Migratate()
	if err != nil {
		fmt.Println(err)
	}

	ech = echo.New()
	route.Route(ech, conn)

	ech.Logger.Fatal(ech.Start(":8016"))
}
