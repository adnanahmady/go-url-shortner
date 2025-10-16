package presentation

import (
	"github.com/adnanahmady/go-url-shortner/pkg/reqeust"
)

type V1Routers struct {
	server     *request.Server
	v1Handlers *V1Handlers
}

func NewV1Routers(
	server *request.Server,
	v1Handlers *V1Handlers,
) *V1Routers {
	return &V1Routers{server: server, v1Handlers: v1Handlers}
}

func (r *V1Routers) Register() {
	r.server.Handle("/", r.v1Handlers.Index)
	r.server.Handle("/shorten", r.v1Handlers.CreateShortUrl)
	r.server.Handle("/{shortUrl}", r.v1Handlers.RedirectToOriginalUrl)
}
