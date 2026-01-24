-- +goose Up
-- +goose StatementBegin
CREATE TYPE AUTH_PROVIDER AS ENUM (
    'GOOGLE',
    'FACEBOOK'
);

CREATE TABLE users (
    id              UUID NOT NULL DEFAULT uuidv7(),
    created_at      TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email           TEXT NOT NULL,
    auth_id         TEXT NOT NULL,
    auth_provider   AUTH_PROVIDER NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (auth_id, auth_provider)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE AUTH_PROVIDER;
-- +goose StatementEnd
