package request

import (
	"github.com/google/uuid"
)

//TODO: add accurate tags
type PostReferrer struct {
	ReferralCode string    `json:"referral_code"`
	ReferredBy   uuid.UUID `json:"referred_by"`
}

type PostTaskComplete struct {
	TaskType string `json:"task_type"`
}
