package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom validator
	e.Validator = &CustomValidator{validator: validator.New()}
	registerHandlers(e)
	err := e.Start(":8090")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
