package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/railgodev/denet-test/internal/apperr"
	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func (uc *usecase) Status(ctx context.Context, id string) (*usecaseModel.Status, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, apperr.ErrInvalidIDFormat
	}
	return uc.repo.Status(ctx, id)
}
