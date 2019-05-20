CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create schema roomba;

create table roomba.rooms (
		id uuid primary key default uuid_generate_v1(),
		room_width integer,
		room_height integer
);

create table roomba.dirt (
		id uuid primary key default uuid_generate_v1(),
		room_id uuid references roomba.rooms(id),
		x_pos integer,
		y_pos integer
);

create table roomba.prior_runs (
		id uuid primary key default uuid_generate_v1(),
		room_id uuid references roomba.rooms(id),
		start_x integer,
		start_y integer,
		instructions text,
		finish_x integer,
		finish_y integer,
		dirt_collected integer
);
		


