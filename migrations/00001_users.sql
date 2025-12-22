-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id              UUID NOT NULL DEFAULT uuidv7(),
    created_at      TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email           TEXT NOT NULL,
    auth_id         TEXT NOT NULL,
    auth_provider   TEXT NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (auth_id, auth_provider)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
