-- +goose Up
-- +goose StatementBegin
CREATE TABLE workout_exercises (
    id              UUID NOT NULL DEFAULT uuidv7(),
    workout_id      UUID NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
    exercise_id     UUID NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,

    position    INTEGER NOT NULL,
    notes       TEXT,

    PRIMARY KEY (id),
    UNIQUE (workout_id, position)
);

CREATE INDEX workout_exercises_workout_id_idx ON workout_exercises(workout_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX workout_exercises_workout_id_idx;
DROP TABLE workout_exercises;
-- +goose StatementEnd
