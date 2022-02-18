package proto_http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"

	"emwell/internal/api/.protobuf/proto"
	"emwell/internal/logger"
)

type Server struct {
	logger   logger.ILogger
	url      string
	handlers proto.EmWellServiceServer
}

func NewServer(logger logger.ILogger, url string, handlers proto.EmWellServiceServer) *Server {
	return &Server{
		logger:   logger,
		url:      url,
		handlers: handlers,
	}
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	httpServer := &http.Server{
		Addr: s.url,
	}

	mux := runtime.NewServeMux()
	httpServer.Handler = mux
	if err := proto.RegisterEmWellServiceHandlerServer(ctx, mux, s.handlers); err != nil {
		return err
	}

	s.logger.InfoKV(ctx, "http server started", "url", s.url)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.ErrorKV(ctx, "Error server.ListenAndServe", zap.Error(err))
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
