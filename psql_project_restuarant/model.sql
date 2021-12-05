create database halal_meals;

drop table if exists components cascade;
drop table if exists categories cascade;
drop table if exists types cascade;
drop table if exists type_register cascade;
drop table if exists tables cascade;
drop table if exists products cascade;
drop table if exists ingredients cascade;
drop table if exists orders cascade;
drop table if exists order_details cascade;

create table components (
  component_id serial not null primary key,
  name varchar(32) not null
);

create table categories (
  category_id serial not null primary key,
  name varchar(64) not null
);

create table types (
  type_id serial not null primary key,
  name varchar(24) not null
);

create table type_register (
  type_register_id serial not null primary key,
  category_id int not null references categories(category_id),
  type_id int not null references types(type_id)
);

create table tables (
  table_id serial not null primary key,
  number int2 not null
);

create table products (
  product_id serial not null primary key,
  name varchar(64) not null,
  price decimal(16, 2) not null,
  type_register_id int not null references type_register(type_register_id)
);

create table ingredients (
  ingredient_id serial not null primary key,
  product_id int not null references products(product_id),
  component_id int not null references components(component_id),
  weight int2 not null
);

create table orders (
  order_id serial not null primary key,
  table_id int not null references tables(table_id),
  created_at timestamp with time zone default current_timestamp,
  closed_at timestamp with time zone default null
);

create table order_details (
  order_detail_id serial not null primary key,
  quantity int2 default 1,
  order_id int not null references orders (order_id),
  product_id int not null references products(product_id)
);

-- COMMENTS

comment on table components is 'Xom ashyolar';
comment on table categories is 'Kategoriyalar';
comment on table types is 'Maxsulot turlari';
comment on table type_register is 'Qaysi kategoriyada qanday turdagi maxsulotlar borligi';
comment on table tables is 'Stollar';
comment on table products is 'Tayyor maxsulotlar';
comment on table ingredients is 'Retseptlar';
comment on table orders is 'Buyurtmalar';
comment on table orders is 'Buyurtma tafsilotlari';
