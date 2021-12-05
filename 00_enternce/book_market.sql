create database book_market;

create extension if not exists "uuid-ossp";

drop table if exists categories cascade;
drop table if exists authors cascade;
drop table if exists books cascade;
drop table if exists clients cascade;
drop table if exists orders cascade;

create table categories (
	category_id serial not null primary key,
	title character varying(32) not null
);

create table authors (
	author_id serial not null primary key,
	fullname character varying(64) not null
);

create table books (
	book_id serial not null primary key,
	category_id int not null references categories (category_id),
	author_id int not null references authors (author_id),
	book_name character varying(32) not null,
	price decimal(8,2) not null
);

create table clients (
	client_id serial not null primary key,
	fullname character varying(65) not null,
	phone_number character varying(9) not null
);

create table orders (
	order_id uuid default uuid_generate_v1 () primary key,
	client_id int not null references clients (client_id),
	book_id int not null references books (book_id),
	created_at timestamptz default current_timestamp
);

\
-- MOCK DATA
insert into categories (title) values ('ilmiy kitoblar'), ('badiiy kitoblar');
insert into authors (fullname) values ('Mansur Usmanov'), ('Tohir Malik');
insert into books (category_id, author_id, book_name, price) values (1, 1, 'Fizika', 24999.99), (2, 2, 'Shayatanat I', 39999.99);
insert into clients (fullname, phone_number) values ('Jasurbek Shamsitdionov', '931062107');
insert into orders (client_id, book_id) values (1, 1), (1, 2);

-- QUERIES

select
	c.category_id,
	c.title,
	a.author_id,
	a.fullname,
	b.book_id,
	b.book_name,
	b.price
from books as b
join categories as c using (category_id)	
join authors as a using (author_id)	
;

