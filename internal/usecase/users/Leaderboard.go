package usecase

import (
	"context"

	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func (uc *usecase) Leaderboard(ctx context.Context) ([]usecaseModel.Leader, error) {
	leaderboard, err := uc.repo.Leaderboard(ctx)
	return leaderboard, err
}
