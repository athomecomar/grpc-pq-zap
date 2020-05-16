CREATE TABLE users (
  id integer unique,
  email   varchar(40)
  password_hash    varchar(40),
);