package main

import (
	"github.com/labstack/echo/v4"

	"github.com/lehuynhthang99/learn_to_do_list_backend/middleware"
)

const (
	certificate_path string = "Key/server.crt"
	private_key_path string = "Key/server.key"
)

func main() {
	e := echo.New()

	//Add middleware
	e.Use(middleware.LoggerMiddleware)
	e.Use(middleware.AuthMiddleWare)

	//Start server
	e.Logger.Fatal(e.StartTLS(":1323", certificate_path, private_key_path))
}
