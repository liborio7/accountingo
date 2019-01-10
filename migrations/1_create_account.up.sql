create table account
(
	id uuid not null
		constraint account_pk
			primary key,
	version integer default 0,
	insert_date timestamp default now(),
	insert_millis bigint default (date_part('epoch'::text, now()) * (1000000)::double precision),
	name text,
	surname text,
	currency text,
	amount bigint default 0
);

create index account_insert_millis_index
	on account (insert_millis desc);

