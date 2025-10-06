package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/railgodev/denet-test/internal/apperr"
)

func (r *repo) TaskCompletion(ctx context.Context, id, taskType string) error {
	// TODO: add atomacity with transaction
	var exists bool
	err := r.conn.QueryRow(ctx,
		`SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)`, id,
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return apperr.ErrUserNotFound
	}
	var TaskCompleted bool
	err = r.conn.QueryRow(ctx,
		`SELECT EXISTS (
			SELECT 1 
			FROM completed_tasks ct
			JOIN task_types tt
			ON ct.task_type_id = tt.id
			WHERE ct.user_id = $1 AND tt.task_type = $2
		)`,
		id,
		taskType,
	).Scan(&TaskCompleted)
	if err != nil {
		return fmt.Errorf("failed to check if task is already completed: %w", err)
	}
	if TaskCompleted {
		return nil
	}

	var TaskPoints int
	err = r.conn.QueryRow(ctx,
		`SELECT points 
		FROM task_types
		WHERE task_type = $1;`,
		taskType,
	).Scan(&TaskPoints)
	if err != nil {
		return fmt.Errorf("failed to check task points: %w", err)
	}
	if err == pgx.ErrNoRows {
		return apperr.ErrTaskNotFound
	}
	_, err = r.conn.Exec(ctx,
		`UPDATE users 
		SET points = points + $2
		WHERE id = $1;`,
		id,
		TaskPoints,
	)

	if err != nil {
		return fmt.Errorf("failed to update user points: %w", err)
	}
	_, err = r.conn.Exec(ctx,
		`INSERT INTO completed_tasks (user_id, task_type_id)
		VALUES (
			$1,
			(SELECT id 
			FROM task_types 
			WHERE task_type = $2)
		)`,
		id,
		taskType,
	)

	if err != nil {
		return fmt.Errorf("failed to insert entry to completed_tasks: %w", err)
	}

	return nil
}
