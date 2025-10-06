-- +goose Down
DROP TRIGGER IF EXISTS trg_set_updated_at ON users;
DROP FUNCTION IF EXISTS set_updated_at();
DROP EXTENSION IF EXISTS pgcrypto;
DROP TABLE IF EXISTS completed_tasks;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS task_types;
