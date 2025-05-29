package API_Service

import (
	"API_Service/internal/config"
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HTTPServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
