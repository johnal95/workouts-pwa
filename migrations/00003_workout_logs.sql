-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_logs (
    id          UUID NOT NULL DEFAULT uuidv7(),
    workout_id  UUID NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
    date        DATE NOT NULL,

    PRIMARY KEY ("id")
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workout_logs;
-- +goose StatementEnd
