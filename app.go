package main

import (
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/datastore"
	"go-boilerplate/server"
)

func main() {
	config.InitConfigs(".env")

	datastore.Connect()
	fmt.Println("Connected to Database successfully")

	server.Init()
	server.InitRouter()
	fmt.Println("Echo server and Router initialized")

	fmt.Println("Starting Echo server")
	server.Connect()
}
