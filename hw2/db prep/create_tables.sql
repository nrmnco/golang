create table faculties (
	id serial primary key,
	name varchar(30) not null
);

create table groups (
	id serial primary key,
	name varchar(30) not null,
	faculty_id integer references faculties(id)
);

create table students (
	id serial primary key,
	first_name varchar(50) not null,
	last_name varchar(50) not null,
	sex varchar(1) check (sex in ('m', 'f')),
	birth_date date,
	group_id integer references groups(id)
);

create table schedule (
	id serial primary key,
	group_id integer references groups(id),
	faculty_id integer references faculties(id),
	day_of_week varchar(3) check (day_of_week in ('mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun')),
	start_time time,
	end_time time
);