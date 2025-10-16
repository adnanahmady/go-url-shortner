package request

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/adnanahmady/go-url-shortner/pkg/applog"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Server struct {
	mux         *http.ServeMux
	middlewares []Middleware
	logger      applog.Logger
	server      *http.Server
}

func NewServer(logger applog.Logger) *Server {
	return &Server{
		mux:         http.NewServeMux(),
		middlewares: []Middleware{},
		logger:      logger,
	}
}

func (s *Server) Use(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, s.middlewares[0](handler))
}

func (s *Server) Run(host string) error {
	s.logger.Info("Starting server", "host", host)
	s.server = &http.Server{
		Addr:    host,
		Handler: s.mux,
	}
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start server (%v): %w", host, err)
	}
	return nil
}

func (s *Server) Shutdown() {
	s.logger.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Errorf("Error shutting down server: %v", err)
	}
	s.logger.Info("Server shut down successfully")
}
