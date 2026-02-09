create table users (
	id serial primary key,
	name varchar(100) not null,
	email varchar(150) unique not null,
	phone varchar(20),
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp
);

insert into users (name, email, phone)
	values ('Sadia', 'sadiajessia@gmail.com', '+8801727392836');


select * from users;