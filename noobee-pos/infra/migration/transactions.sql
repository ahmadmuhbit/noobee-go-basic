create table transactions(
	id serial primary key,
	employee_id int,
	menu_id int,
	quantity int,
	total int,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

alter table transactions 
add foreign key (menu_id) references menus (id);

alter table transactions 
add foreign key (employee_id) references employees (id);