-- +goose Up
-- +goose StatementBegin
CREATE TABLE translations (
    id SERIAL PRIMARY KEY,
    entity_type VARCHAR(50) NOT NULL,
    entity_id INTEGER NOT NULL,
    language VARCHAR(10) NOT NULL,
    field_name VARCHAR(50) NOT NULL,
    translation TEXT NOT NULL,
    UNIQUE(entity_type, entity_id, language, field_name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE translations;
-- +goose StatementEnd
