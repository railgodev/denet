package usecase

import (
	"log/slog"

	"github.com/railgodev/denet-test/internal/repo"
)

type usecase struct {
	repo repo.Users

	log *slog.Logger
}

func New(repo repo.Users, log *slog.Logger) *usecase {
	return &usecase{
		repo: repo,
		log:  log,
	}
}
