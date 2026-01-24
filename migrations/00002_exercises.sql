-- +goose Up
-- +goose StatementBegin
CREATE TYPE EXERCISE_TYPE AS ENUM (
    'REPS',
    'REPS_WEIGHT',
    'DURATION'
);

CREATE TABLE exercises (
    id          UUID NOT NULL DEFAULT uuidv7(),
    created_at  TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    type EXERCISE_TYPE NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE exercise_translations (
    exercise_id     UUID NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    locale          TEXT NOT NULL,
    name            TEXT NOT NULL,
    description     TEXT,

    PRIMARY KEY (exercise_id, locale),
    UNIQUE (locale, name)
);

CREATE INDEX exercise_translations_exercise_id_idx ON exercise_translations(exercise_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE exercises;
DROP INDEX exercise_translations_exercise_id_idx;
DROP TABLE exercise_translations;
DROP TYPE EXERCISE_TYPE
-- +goose StatementEnd
