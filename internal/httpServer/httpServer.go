package httpServer

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/denet-test/internal/config"
)

type Server struct {
	httpsrv *http.Server
	logger  *slog.Logger
}

func New(cfg *config.Config, handler *gin.Engine, logger *slog.Logger) *Server {
	httpsrv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTP.Port),
		Handler: handler,
	}
	return &Server{
		httpsrv: httpsrv,
		logger:  logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting http server", "addr", s.httpsrv.Addr)
	if err := s.httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpsrv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	s.logger.Info("server stopped")
	return nil
}
