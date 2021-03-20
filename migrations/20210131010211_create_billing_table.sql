-- +goose Up
-- +goose StatementBegin

CREATE TABLE "billing" (
     "id"         BIGSERIAL PRIMARY KEY,
     "user_id"       INTEGER NOT NULL,
     "identify"       TEXT NOT NULL,
     "key"   TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS billing;
-- +goose StatementEnd
