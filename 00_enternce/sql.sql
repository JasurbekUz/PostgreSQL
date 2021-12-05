-- SQL - Structure Query Language  (case ins)

-- SQL Products: PostgreSQL, MS_SQL, MySQL, SQLite, Oracle


	-- DDL - Data Definition Langugage (create, drop ...)
		-- create [kind] [name]
		-- drop [kind] [name]
		--ex:
			create database library;

			create table books(
				book_id serial not null primary key,
				book_name character varying(64)
			);

			drop table books;

			-- relation model (table - relation)

	-- DML - Data Manipulation Laguage (insert)
		--ex:
			insert into books (book_name) values ('book name');

	-- DQL - Data Query Laguage (select, ...)
	
			select * from library;

			select 'hello, world!';

			select 3<4;
			select 5+2;

			select now();
			select extract(year from now());