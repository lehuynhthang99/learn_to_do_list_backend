package main

import (
	"github.com/lehuynhthang99/learn_to_do_list_backend/config"
	"github.com/lehuynhthang99/learn_to_do_list_backend/routes"
)

const (
	certificate_path string = "Key/server.crt"
	private_key_path string = "Key/server.key"
)

func main() {
	config.ConnectDatabase()

	e := routes.RegisterRoutes()

	//Start server
	e.Logger.Fatal(e.StartTLS(":1323", certificate_path, private_key_path))
}
