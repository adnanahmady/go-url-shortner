package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adnanahmady/go-url-shortner/internal"
)

const (
	storageFile = "urls.json"
	port        = ":5000"
)

func main() {
	app, err := internal.InitializeServer()
	if err != nil {
		log.Fatal(err)
	}
	loadUrls(app)
	setupGracefulShutdown(app)

	app.Server.Use(app.LoggingMiddleware.Middleware)

	app.V1Routers.Register()

	app.Server.Run(port)
	app.Logger.Infof("Server stoped on port (%v)", port)
}

func loadUrls(app *internal.App) {
	app.StoreManager.Lock()
	defer app.StoreManager.Unlock()

	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			app.Logger.Infof("Storage file (%v) not found, starting with empty store.", storageFile)
			return
		}

		app.Logger.Errorf("Error reading storage file (%v): %v", storageFile, err)
		os.Exit(1)
	}

	err = app.StoreManager.LoadFromJSON(data)
	if err != nil {
		app.Logger.Errorf("Failed to unmarshal URLs from storage file (%v): %v", storageFile, err)
		os.Exit(1)
	}

	app.Logger.Infof(
		"Storage file (%v) loaded (%v) URLs successfully.",
		storageFile,
		app.StoreManager.Count(),
	)
}

func setupGracefulShutdown(app *internal.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		app.Logger.Infof("Starting graceful shutdown...")
		saveUrls(app)
		app.Logger.Infof("Graceful shutdown completed successfully")
		os.Exit(0)
	}()
}

func saveUrls(app *internal.App) {
	app.StoreManager.Lock()
	defer app.StoreManager.Unlock()

	data, err := app.StoreManager.ToJSON()
	if err != nil {
		app.Logger.Errorf("Failed to get Store as JSON: %v", err)
		return
	}

	if err := os.WriteFile(storageFile, data, 0644); err != nil {
		app.Logger.Errorf("Failed to write URLs to storage file: %v", err)
		return
	}

	app.Logger.Infof(
		"%v URLs saved to storage file (%v) successfully.",
		app.StoreManager.Count(),
		storageFile,
	)
}
