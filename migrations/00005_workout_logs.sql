-- +goose Up
-- +goose StatementBegin
CREATE TABLE workout_logs (
    id          UUID NOT NULL DEFAULT uuidv7(),
    workout_id  UUID NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
    date        DATE NOT NULL,

    PRIMARY KEY (id)
);

CREATE INDEX workout_logs_workout_id_idx ON workout_logs(workout_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX workout_logs_workout_id_idx;
DROP TABLE workout_logs;
-- +goose StatementEnd
