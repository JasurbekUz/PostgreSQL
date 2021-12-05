create table authors (
	author_id serial not null primary key,
	fullname character varying(64) not null,
	nickname character varying(16),
	birth date not null
);

create table categories (
	category_id serial not null primary key,
	category_name character varying(32) not null
);

create table books (
	book_id serial not null primary key,
	category_id int not null references categories (category_id),
	author_id int not null references authors (author_id),
	book_name character varying(32) not null
);