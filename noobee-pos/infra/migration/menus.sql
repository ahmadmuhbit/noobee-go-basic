create table menus(
	id serial primary key,
	name varchar(100) not null,
	category varchar(100) not null,
	description varchar(255),
	price int not null,
	image_url varchar(255) not null,
	created_at timestamp default now(),
	updated_at timestamp default now()
);