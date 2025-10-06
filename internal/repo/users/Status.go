package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/railgodev/denet-test/internal/apperr"
	repoConverter "github.com/railgodev/denet-test/internal/repo/converter"
	repoModel "github.com/railgodev/denet-test/internal/repo/model"
	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func (r *repo) Status(ctx context.Context, id string) (*usecaseModel.Status, error) {
	var status repoModel.Status

	err := r.conn.QueryRow(ctx,
		`SELECT u.id, u.username, u.email, u.points, u.referral_code, ref.username AS referred_by, u.created_at
		FROM users u
		LEFT JOIN users ref ON u.referred_by = ref.id
		WHERE u.id = $1;`,
		id,
	).Scan(
		&status.ID,
		&status.Username,
		&status.Email,
		&status.Points,
		&status.ReferralCode,
		&status.ReferredBy,
		&status.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrUserNotFound
		}
		return nil, err
	}
	return repoConverter.StatusToUsecaseModel(&status), nil
}
