package usecase

import (
	"context"

	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func (uc *usecase) Referrer(ctx context.Context, id string, referrer *usecaseModel.Referrer) error {
	return (uc.repo.Referrer(ctx, id, referrer))
}
