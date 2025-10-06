package handlerConverter

import (
	"github.com/railgodev/denet-test/internal/api/v1/request"
	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func ReferrerToRepoModel(referrer *request.PostReferrer) *usecaseModel.Referrer {
	return &usecaseModel.Referrer{
		ReferralCode: referrer.ReferralCode,
		ReferredBy:   referrer.ReferredBy,
	}
}
