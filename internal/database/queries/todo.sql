-- name: CreateTodo :one
INSERT INTO "todo" (name, user_id) VALUES ($1, $2) RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM "todo" WHERE id = $1;

-- name: GetTodoById :one
SELECT * FROM "todo" WHERE id = $1;

-- name: GetTodos :many
SELECT * FROM "todo" WHERE user_id = $1 ORDER BY id;

-- name: UpdateTodoCompleted :one
UPDATE "todo" SET completed = $2 WHERE id = $1 RETURNING *;

-- name: UpdateTodoName :one
UPDATE "todo" SET name = $2 WHERE id = $1 RETURNING *;
