-- +goose Up
-- +goose StatementBegin
CREATE TABLE emotional_diary
(
    id             BIGSERIAL PRIMARY KEY,
    user_id        BIGINT REFERENCES users (id),
    emotion_rate   INTEGER                  NOT NULL,
    refers_to_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE emotional_diary;
-- +goose StatementEnd
