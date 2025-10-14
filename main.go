package main

import (
	"log"

	"github.com/adnanahmady/go-url-shortner/internal"
)

func main() {
	app, err := internal.InitializeServer()
	if err != nil {
		log.Fatal(err)
	}

	app.Server.Use(app.LoggingMiddleware.Middleware)

	app.V1Routers.Register()

	app.Server.Run(":5000")
}