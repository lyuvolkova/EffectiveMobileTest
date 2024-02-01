-- +goose Up
-- +goose StatementBegin
create table person
(
    user_id       serial             not null
        constraint user_pk
            primary key,
    name          text,
    surname       text,
    patronymic    text,
    age           integer,
    gender        integer,
    nationalize   text,
    created_at    timestamptz default NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table person;
-- +goose StatementEnd
