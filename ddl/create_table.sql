CREATE TABLE tasks(
	id serial primary key,
	content text not null,
	is_completed boolean not null
);
