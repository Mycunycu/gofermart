CREATE TABLE IF NOT EXISTS customer(
   id serial PRIMARY KEY,
   login VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (128) NOT NULL
);