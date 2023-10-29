insert into menus(name, category, description, price, image_url)
values('Ketoprak', 'Makanan', 'Ketoprak Cirebon', 14000, 'https://rebrand.ly/h6k9u3x');

insert into menus(name, category, description, price, image_url)
values('Nasi Goreng', 'Makanan', 'Nasi Goreng Spesial', 18000, 'https://unsplash.com/photos/rQX9eVpSFz8');

insert into menus(name, category, description, price , image_url)
values('Jus Alpukat', 'Minuman', 'Jus Alpukat Susu', 10000, 'https://unsplash.com/photos/_82CV9I-TP8'),
		('Es Teh Manis', 'Minuman', 'Es Teh Manis', 4000, 'https://unsplash.com/photos/CCowelQ2pLw');

	
select * from menus;

select name, category, price, image_url
from menus;

select * from menus
where id = 1;

update menus 
set name = 'Ketoprak Edit', price = 12000
where id = 1;

delete from menus
where id = 3;

select * from employees;

insert into employees(nip, name, address)
values('00001', 'Noobee Id', 'Bintaro, Tangerang Selatan'),
		('00002', 'Ahmad Muhbit', 'Kelapa Gading, Jakarta Selatan');	
		
select * from transactions;

insert into transactions(menu_id, employee_id, quantity, total)
values(1, 1, 2, 24000),	
		(2, 2, 1, 18000);

select t.id, t.employee_id, t.menu_id, t.quantity, t.total,
		m.id, m.name, m.price,
		e.id, e.name
		from transactions as t
		join employees as e on e.id = t.employee_id 
		join menus as m on m.id = t.menu_id;
		