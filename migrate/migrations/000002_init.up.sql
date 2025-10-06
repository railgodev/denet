-- +goose Up
INSERT INTO
  task_types (task_type, points)
VALUES
  ('sub on Yandex Plus', 1000),
  ('sub on Netflix', 1000),
  ('sub on premium', 1000);
INSERT INTO
  users (id, username, email, password_hash)
VALUES
  (
    'dd74ba03-5608-4ad2-b251-b0e87843d0a1',
    '1',
    'u1@gmail.com',
    'pass_hash1'
  ),
  (
    'dd74ba03-5608-4ad2-b251-b0e87843d0a2',
    '2',
    'u2@gmail.com',
    'pass_hash2'
  );