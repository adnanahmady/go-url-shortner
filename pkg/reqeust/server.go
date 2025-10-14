package request

import (
	"log"
	"net/http"

	"github.com/adnanahmady/go-url-shortner/pkg/applog"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Server struct {
	mux *http.ServeMux
	middlewares []Middleware
	logger applog.Logger
}

func NewServer(logger applog.Logger) *Server {
	return &Server{
		mux: http.NewServeMux(),
		middlewares: []Middleware{},
		logger: logger,
	}
}

func (s *Server) Use(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, s.middlewares[0](handler))
}

func (s *Server) Run(host string) {
	s.logger.Info("Starting server", "host", host)
	log.Fatal(http.ListenAndServe(host, s.mux))
}
