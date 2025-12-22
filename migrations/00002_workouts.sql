-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workouts (
    id          UUID NOT NULL DEFAULT uuidv7(),
    created_at  TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id     UUID NOT NULL REFERENCES users(id),
    name        TEXT NOT NULL,

    PRIMARY KEY ("id"),
    UNIQUE (user_id, name)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workouts;
-- +goose StatementEnd
