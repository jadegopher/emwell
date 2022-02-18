package http

import (
	"context"
	"net/http"
	"time"

	"emwell/internal/logger"
)

type Server struct {
	logger   logger.ILogger
	url      string
	handlers http.Handler
}

func NewServer(logger logger.ILogger, url string, handlers http.Handler) *Server {
	return &Server{
		logger:   logger,
		url:      url,
		handlers: handlers,
	}
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	httpServer := &http.Server{
		Addr:         s.url,
		Handler:      s.handlers,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		s.logger.InfoKV(ctx, "http server started", "url", s.url)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.ErrorKV(ctx, "Error server.ListenAndServe", "err", err)
			return
		}
	}()

	for range ctx.Done() {
		if err := httpServer.Shutdown(ctx); err != nil {
			s.logger.ErrorKV(ctx, "http server shutdown error", "err", err)
		}
		s.logger.InfoKV(ctx, "http server exited properly")
		break
	}

	return nil
}
