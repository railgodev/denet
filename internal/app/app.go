package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/denet-test/internal/api"
	"github.com/railgodev/denet-test/internal/config"
	"github.com/railgodev/denet-test/internal/httpServer"
	"github.com/railgodev/denet-test/internal/logger"
	repo "github.com/railgodev/denet-test/internal/repo/users"
	"github.com/railgodev/denet-test/internal/storage"
	usecase "github.com/railgodev/denet-test/internal/usecase/users"
	"github.com/railgodev/denet-test/migrate"

	l "log"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	log := logger.New(cfg.Log.Level)
	log.Debug("config", slog.Any("config", cfg))

	if err := migrate.Run(cfg.PG.URL, cfg.Migrations.MigratePath, log); err != nil {
		l.Fatal(fmt.Errorf("app - migrate - Run: %w", err))
	}
	conn, err := storage.GetConn(log, cfg.PG.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - storage.GetConn: %w", err))
	}
	defer conn.Close()
	repo := repo.New(conn, log)

	uc := usecase.New(repo, log)

	if cfg.Log.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	api := api.New(cfg, uc, log)

	httpServer := httpServer.New(cfg, api, log)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpServer.Start(ctx); err != nil {
			log.Error("failed to start HTTP server", slog.Any("err", err))
		}
	}()

	<-done

	if err := httpServer.Stop(); err != nil {
		log.Error("failed to stop HTTP server", slog.Any("err", err))
	}
}
