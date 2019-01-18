package server

import (
	"github.com/labstack/echo"
	"go-boilerplate/config"
	"log"
	"net/http"
)

/*
We keep the echo server instance as a global
variable to be used across packages
*/
var EchoCon *echo.Echo

func Init() {
	EchoCon = echo.New()
	http.Handle("/", EchoCon)
}

/*
Calling this method will block the main thread 
until further callbacks from the echo framework
*/
func Connect() {
	if EchoCon == nil {
		log.Fatal("Echo Server has not been initialized. Instance is nil")
		return
	}
	
	EchoCon.Logger.Fatal(EchoCon.Start(":" + config.GetConfig("SERVER_PORT")))
}





