package test

import (
	"github.com/adnanahmady/go-url-shortner/internal"
)

func Setup() (*internal.App, error) {
	app, err := internal.InitializeServer()
	if err != nil {
		return nil, err
	}

	app.Server.Use(app.LoggingMiddleware.Middleware)

	app.V1Routers.Register()

	return app, nil
}
