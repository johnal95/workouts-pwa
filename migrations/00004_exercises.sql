-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS exercises (
    id                  UUID NOT NULL DEFAULT uuidv7(),
    created_at          TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    workout_id          UUID NOT NULL REFERENCES workouts(id),
    name                TEXT NOT NULL,
    default_set_count   INTEGER NOT NULL,
    min_reps            INTEGER NOT NULL,
    max_reps            INTEGER NOT NULL,

    PRIMARY KEY ("id"),
    UNIQUE (workout_id, name)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE exercises;
-- +goose StatementEnd
