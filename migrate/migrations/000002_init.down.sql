-- +goose Down
DELETE FROM task_types
WHERE
  task_type IN (
    'sub on Yandex Plus',
    'sub on Netflix',
    'sub on premium'
  );

DELETE FROM users
WHERE
  id IN (
    'dd74ba03-5608-4ad2-b251-b0e87843d0a1',
    'dd74ba03-5608-4ad2-b251-b0e87843d0a2'
  );