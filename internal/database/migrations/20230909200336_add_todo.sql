-- +goose Up
-- +goose StatementBegin
CREATE TABLE "todo" (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  user_id INTEGER NOT NULL REFERENCES "user"(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "todo";
-- +goose StatementEnd
