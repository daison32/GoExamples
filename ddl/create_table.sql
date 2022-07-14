CREATE TABLE tasks(
	id serial primary key,
	content text not null,
	is_completed boolean not null
);

ALTER TABLE tasks ALTER COLUMN is_completed boolean not null SET DEFAULT false;

