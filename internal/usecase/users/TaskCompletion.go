package usecase

import "context"

func (uc *usecase) TaskCompletion(ctx context.Context, id, taskType string) error {
	return uc.repo.TaskCompletion(ctx, id, taskType)
}
