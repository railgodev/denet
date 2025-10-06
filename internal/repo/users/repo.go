package repo

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	conn *pgxpool.Pool
	log  *slog.Logger
}

func New(conn *pgxpool.Pool, log *slog.Logger) *repo {
	return &repo{conn: conn, log: log}
}

func (r *repo) Close() {
	r.conn.Close()
}
