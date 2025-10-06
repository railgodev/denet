package repoModel

import (
	"time"

	"github.com/google/uuid"
)

type Leader struct {
	ID       string
	Username string
	Email    string
	Points   int
}

type Status struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Points       int
	ReferralCode *string
	ReferredBy   *uuid.UUID
	CreatedAt    time.Time
}

type Referrer struct {
	ReferralCode string
	ReferredBy   uuid.UUID
}
