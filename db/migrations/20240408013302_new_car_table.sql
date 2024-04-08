-- +goose Up
-- +goose StatementBegin
create table if not exists car
(regNum varchar primary key ,
 mark varchar, model varchar,
 year integer, owner integer);

create table if not exists people
(id varchar primary key,
 name varchar , surname varchar,
 patronymic varchar);

create table if not exists nums
(regNum varchar primary key,
 owner_id varchar);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS car, people;
-- +goose StatementEnd
