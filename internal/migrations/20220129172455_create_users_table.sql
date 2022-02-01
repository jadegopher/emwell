-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(250)             NOT NULL,
    language    VARCHAR(3)               NOT NULL,
    telegram_id BIGINT                   NOT NULL UNIQUE,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
