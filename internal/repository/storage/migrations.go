package storage

const createTablePhotosTemplate = `create table if not exists Photos (
    id serial primary key,
    name text not null,
    created timestamp not null,
    edited timestamp not null,
    file_path text not null
)`
