create database green_life;

drop table if exists categories cascade;
drop table if exists plants cascade;

create table categories (
	category_id serial not null primary key,
	name character varying(32) not null
);

create table plants (
	plant_id serial not null primary key,
	category_id int not null references categories (category_id),
	plant_name character varying(32) not null,
	quantity int
);

-- MOCK DATA
insert into categories (name) values ('manzarali uy o''simliklari'), ('manzarali daraxtlar'), ('sovg''a uchun gullar');
insert into plants (category_id, plant_name, quantity) values (1, 'limon', 30), (3, 'qizil lola', 50), (2, 'sasna', 12);