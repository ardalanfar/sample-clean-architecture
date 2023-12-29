
create database farashop;

CREATE TABLE IF NOT EXISTS "user" (
id serial primary key,
usrename varchar(50) unique not null,
password varchar(50) not null,
email varchar(255) unique not null,
access int unique not null,
address varchar(400),
Verification_code int not null,
Is_verified varchar(50) not null
);

CREATE TABLE IF NOT EXISTS "wallet" (
id serial primary key,
user_id int not null,
balance int not null
);

CREATE TABLE IF NOT EXISTS "product" (
id serial primary key,
name varchar(255) unique not null,
price int not null,
number int not null,
describe text
);

CREATE TABLE IF NOT EXISTS "order" (
id serial primary key,
user_id int not null,
product_id int not null,
status int not null,
number int not null,
buy_time TIMESTAMP
);
