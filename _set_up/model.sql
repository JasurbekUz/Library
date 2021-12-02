--create database library;

drop table if exists authors cascade;
drop table if exists categories cascade;
drop table if exists books cascade;

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

--MOCK DATA
	-- authors
insert into authors (fullname, nickname, birth) values ('Alisher Navoiy', 'Navoiy', '09-02-1441');
insert into authors (fullname, nickname, birth) values ('Musa Toshmuhammad o''g''li', 'Oybek', '10-01-1905');

	-- categories
insert into categories (category_name) values ('mumtoz'), ('tarixiy');

	-- books
insert into books (category_id, author_id, book_name) values (1, 1, 'Hamsa'), (2, 2, 'Navoiy');

--QUERIES
select 
	b.book_name,
	c.category_name,
	a.fullname,
	a.nickname
from
	books as b
join categories c using (category_id)
join authors a using (author_id)
;
