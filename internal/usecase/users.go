package usecase

import (
	"context"

	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

type (
	Users interface {
		Leaderboard(ctx context.Context) ([]usecaseModel.Leader, error)
		Status(ctx context.Context, id string) (*usecaseModel.Status, error)
		Referrer(ctx context.Context, id string, referrer *usecaseModel.Referrer) error
		TaskCompletion(ctx context.Context, id, taskType string) error
	}
)
