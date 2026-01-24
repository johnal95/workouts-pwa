-- +goose Up
-- +goose StatementBegin
CREATE TABLE workouts (
    id          UUID NOT NULL DEFAULT uuidv7(),
    created_at  TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name        TEXT NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (user_id, name)
);

CREATE INDEX workouts_user_id_idx ON workouts(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX workouts_user_id_idx;
DROP TABLE workouts;
-- +goose StatementEnd
