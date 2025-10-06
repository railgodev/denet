package repo

import (
	"context"
	"fmt"

	"github.com/railgodev/denet-test/internal/apperr"
	repoConverter "github.com/railgodev/denet-test/internal/repo/converter"
	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func (r *repo) Referrer(ctx context.Context, id string, referrer *usecaseModel.Referrer) error {
	Referrer := repoConverter.ReferrerToRepoModel(referrer)
	var userExists, refExists bool

	err := r.conn.QueryRow(ctx,
		`SELECT
        EXISTS (SELECT 1 FROM users WHERE id = $1) AS user_exists,
        EXISTS (SELECT 1 FROM users WHERE id = $2) AS ref_exists`,
		id,
		Referrer.ReferredBy,
	).Scan(
		&userExists,
		&refExists,
	)

	switch {
	case !userExists:
		return apperr.ErrUserNotFound
	case !refExists:
		return apperr.ErrReferrerNotFound
	case err != nil:
		return fmt.Errorf("failed to check user/referrer existence: %w", err)
	}

	_, err = r.conn.Exec(ctx,
		`UPDATE users 
		SET referral_code = $2, referred_by = $3
		WHERE id = $1;`,
		id,
		Referrer.ReferralCode,
		Referrer.ReferredBy,
	)
	if err != nil {
		return fmt.Errorf("failed to update referrer: %w", err)
	}
	return nil

}
