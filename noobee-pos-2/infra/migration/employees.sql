create table employees(
	id serial primary key,
	nip varchar(50) unique not null,
	name varchar(100) not null,
	address varchar(255),
	created_at timestamp default now(),
	updated_at timestamp default now()
);