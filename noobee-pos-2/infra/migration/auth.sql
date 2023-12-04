create table auth(
    id serial primary key,
    email varchar(100) not null,
    password varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
)