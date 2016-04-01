create table user(
	id int auto_increment primary key,
	name char(12) not null unique,
	passwd char(64) not null
)default charset=utf8;

create table problem(
	id int auto_increment primary key,
	user_id int,
	content text
)default charset=utf8;

create table solution(
	id int auto_increment primary key,
	user_id int not null,
	problem_id char(100) not null,
	content text,
	created datetime
)default charset=utf8;

create table remark(
	id int auto_increment primary key,
	user_id int not null,
	solution_id int not null,
	content text,
	created datetime
)default charset=utf8;

alter table solution add result text;


create table puzzle(
	id int auto_increment primary key,
	user_id int not null,
	title text,
	descr text,
	func_name text,
	content text,
	args_type text,
	rets_type text,
	test_cases text,
	online char(1)
)default charset=utf8;