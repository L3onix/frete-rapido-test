create table if not exists carrier (
	id serial primary key,
  name varchar (100),
  service varchar (100),
  deadline integer,
  price float4,
  dispatcher_id varchar(50)
);
