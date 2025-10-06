package repoConverter

import (
	repoModel "github.com/railgodev/denet-test/internal/repo/model"
	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func LeadersToUsecaseModel(leaders []repoModel.Leader) []usecaseModel.Leader {
	leadersCopy := make([]usecaseModel.Leader, len(leaders))
	for _, l := range leaders {
		lCopy := usecaseModel.Leader{
			ID:       l.ID,
			Username: l.Username,
			Email:    l.Email,
			Points:   l.Points,
		}
		leadersCopy = append(leadersCopy, lCopy)
	}
	return leadersCopy
}

func StatusToUsecaseModel(status *repoModel.Status) *usecaseModel.Status {
	if status == nil {
		return nil
	}
	return &usecaseModel.Status{
		ID:           status.ID,
		Username:     status.Username,
		Email:        status.Email,
		Points:       status.Points,
		ReferralCode: status.ReferralCode,
		ReferredBy:   status.ReferredBy,
		CreatedAt:    status.CreatedAt,
	}
}

func ReferrerToRepoModel(referrer *usecaseModel.Referrer) *repoModel.Referrer {
	return &repoModel.Referrer{
		ReferralCode: referrer.ReferralCode,
		ReferredBy:   referrer.ReferredBy,
	}
}
