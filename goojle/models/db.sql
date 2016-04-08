create table user(
	id int auto_increment primary key,
	name char(12) not null unique,
	passwd char(64) not null
)default charset=utf8;

insert into user values(`1`,"shaalx","shaalx");
insert into user values(`4`,"Anonymous","Anonymous");

create table solution(
	id int auto_increment primary key,
	user_id int not null,
	puzzle_id int not null,
	content text,
	result int,
	created datetime
)default charset=utf8;

create table remark(
	id int auto_increment primary key,
	user_id int not null,
	puzzle_id int not null,
	solution_id int not null,
	content text,
	created datetime
)default charset=utf8;


create table puzzle(
	id int auto_increment primary key,
	user_id int not null,
	title text not null,
	descr text,
	func_name text not null,
	content text,
	args_type text not null,
	rets_type text not null,
	test_cases text not null,
	online char(1)
)default charset=utf8;

create table role(
	id int auto_increment primary key,
	title text not null
)default charset=utf8;


create table result(
	id int auto_increment primary key,
	state text not null,
	run_cost_time text not null,
	content text not null,
	test_case text,
	run_result text,
	error_info text
)default charset=utf8;

create unique index user_index on user(id);
create unique index puzzle_index on puzzle(id);
create unique index result_index on result(id);
create unique index solution_index on solution(id);
create unique index remark_index on remark(id);

CREATE TABLE `session` (
	`session_key` char(64) NOT NULL,
	`session_data` blob,
	`session_expiry` int(11) unsigned NOT NULL,
	PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;