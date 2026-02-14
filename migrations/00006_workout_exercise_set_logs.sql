-- +goose Up
-- +goose StatementBegin
CREATE TABLE workout_exercise_set_logs (
    id                  UUID NOT NULL DEFAULT uuidv7(),
    workout_log_id      UUID NOT NULL REFERENCES workout_logs(id) ON DELETE CASCADE,
    workout_exercise_id UUID NOT NULL REFERENCES workout_exercises(id) ON DELETE CASCADE,

    position    INTEGER NOT NULL,

    reps                INTEGER,
    weight_kg           NUMERIC(6,2),
    duration_seconds    INTEGER,

    CHECK (
        -- EXERCISE_TYPE: 'REPS'
        (reps IS NOT NULL AND weight_kg IS NULL AND duration_seconds IS NULL)
        -- EXERCISE_TYPE: 'REPS_WEIGHT'
     OR (reps IS NOT NULL AND weight_kg IS NOT NULL AND duration_seconds IS NULL)
        -- EXERCISE_TYPE: 'DURATION'
     OR (reps IS NULL AND weight_kg IS NULL AND duration_seconds IS NOT NULL)
    ),

    PRIMARY KEY (id),
    UNIQUE (workout_log_id, workout_exercise_id, position)
);

CREATE INDEX workout_exercise_set_logs_workout_log_id_idx ON workout_exercise_set_logs(workout_log_id);

CREATE INDEX workout_exercise_set_logs_workout_exercise_id_idx ON workout_exercise_set_logs(workout_exercise_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX workout_exercise_set_logs_workout_exercise_id_idx;
DROP INDEX workout_exercise_set_logs_workout_log_id_idx;
DROP TABLE workout_exercise_set_logs;
-- +goose StatementEnd
