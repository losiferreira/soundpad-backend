CREATE TABLE users (
      id serial PRIMARY KEY,
      email VARCHAR (100) UNIQUE NOT NULL,
      name VARCHAR (100) NOT NULL,
);