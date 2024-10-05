-- +goose Up
-- +goose StatementBegin
create table test_table (
	id serial primary key,
	name text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
