package repo

import (
	"context"
	"fmt"

	repoConverter "github.com/railgodev/denet-test/internal/repo/converter"
	repoModel "github.com/railgodev/denet-test/internal/repo/model"
	usecaseModel "github.com/railgodev/denet-test/internal/usecase/model"
)

func (r *repo) Leaderboard(ctx context.Context) ([]usecaseModel.Leader, error) {
	rows, err := r.conn.Query(ctx,
		`SELECT u.id, u.username, u.email, u.points
		FROM users u
		ORDER BY u.points DESC, u.username ASC
		LIMIT 10;`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query leaderboard: %w", err)
	}
	defer rows.Close()

	var leaders []repoModel.Leader
	for rows.Next() {
		var leader repoModel.Leader
		if err := rows.Scan(&leader.ID, &leader.Username, &leader.Email, &leader.Points); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		leaders = append(leaders, leader)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return repoConverter.LeadersToUsecaseModel(leaders), nil
}
