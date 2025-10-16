//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/adnanahmady/go-url-shortner/internal/application"
	"github.com/adnanahmady/go-url-shortner/internal/domain"
	"github.com/adnanahmady/go-url-shortner/internal/infra"
	"github.com/adnanahmady/go-url-shortner/internal/presentation"
	"github.com/adnanahmady/go-url-shortner/pkg/applog"
	"github.com/adnanahmady/go-url-shortner/pkg/reqeust"
	"github.com/adnanahmady/go-url-shortner/pkg/store"
	"github.com/google/wire"
)

type App struct {
	Server            *request.Server
	Logger            applog.Logger
	LoggingMiddleware *request.LoggingMiddleware
	V1Routers         *presentation.V1Routers
	V1Handlers        *presentation.V1Handlers
	StoreManager      store.StoreManager
}

var AppSet = wire.NewSet(
	applog.NewWriter,
	applog.NewApplicationLogger,
	wire.Bind(new(applog.Logger), new(*applog.ApplicationLogger)),

	store.NewMemoryStore,
	wire.Bind(new(store.StoreManager), new(*store.MemoryStoreManager)),

	request.NewServer,
	request.NewLoggingMiddleware,

	presentation.NewV1Routers,
	presentation.NewV1Handlers,

	infra.NewMemoryUrlRepository,
	wire.Bind(new(domain.UrlRepository), new(*infra.MemoryUrlRepository)),

	application.NewCreateShortUrlUseCaseImpl,
	wire.Bind(new(application.CreateShortUrlUseCase), new(*application.CreateShortUrlUseCaseImpl)),
	application.NewGetShortUrlUseCaseImpl,
	wire.Bind(new(application.GetShortUrlUseCase), new(*application.GetShortUrlUseCaseImpl)),

	wire.Struct(new(App), "*"),
)

func InitializeServer() (*App, error) {
	wire.Build(AppSet)
	return nil, nil
}
