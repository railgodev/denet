// Package v1 implements routing paths. Each services in own file.
package api

import (
	"net/http"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/denet-test/internal/config"
	v1 "github.com/railgodev/denet-test/internal/api/v1"
	"github.com/railgodev/denet-test/internal/usecase"
)

func New(cfg *config.Config, uc usecase.Users, log *slog.Logger) *gin.Engine {
	app := gin.Default()

	// Probe
	app.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Routers
	apiV1Group := app.Group("/api/v1/")
	handler := v1.New(uc, log)
	v1.NewUsersRoutes(apiV1Group, handler, cfg)
	return app
}
