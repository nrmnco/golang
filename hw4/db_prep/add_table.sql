create table attendance (
	id serial primary key,
	schedule_id integer references schedule(id),
	student_id integer references students(id),
	date_time timestamp,
	attendance bool default false
);