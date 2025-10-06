package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	middleware "github.com/railgodev/denet-test/internal/api/v1/middleware/auth"
	"github.com/railgodev/denet-test/internal/config"
	"github.com/railgodev/denet-test/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
	uc  usecase.Users
	log *slog.Logger
}

func New(uc usecase.Users, log *slog.Logger) *handler {
	return &handler{
		uc:  uc,
		log: log,
	}
}

func NewUsersRoutes(apiV1Group *gin.RouterGroup, h *handler, cfg *config.Config) {
	if cfg.Swagger.Enabled {
		EnableSwagger(apiV1Group)
	}
	
	users := apiV1Group.Group("/users")
	users.GET("/leaderboard", h.GetLeaderboard)
	// auth middleware on HMAC (symmetric) JWT
	users.Use(middleware.NewHMACMiddleware([]byte(cfg.JWT.SECRET), h.log))
	users.GET("/:id/status", h.GetStatus)
	users.POST("/:id/task/complete", h.PostTaskComplete)
	users.POST("/:id/referrer", h.PostReferrer)
}

func EnableSwagger(apiV1Group *gin.RouterGroup) {

	apiV1Group.StaticFile("/docs/swagger.yaml", "./docs/swagger.yaml")
	apiV1Group.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/api/v1/docs/swagger.yaml"),
	))
}
