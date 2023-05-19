-- CREATE DATABASE
-- CREATE DATABASE fitness_app;

-- Establishes a new connection
\connect fitness_app

-- CREATE SCHEMA \
CREATE SCHEMA private;

-- CREATE ROLE
CREATE ROLE admin WITH LOGIN PASSWORD 'P@ssw0rd';

-- GRANT ALL PRIVILEGES
GRANT ALL PRIVILEGES ON SCHEMA private TO admin;

-- CREATE TABLE | private.users
CREATE TABLE private.users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- CREATE TABLE | private.measurements
CREATE TABLE private.measurements (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  weight FLOAT NOT NULL,
  height FLOAT NOT NULL,
  body_fat FLOAT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES private.users(id)
);


-- CREATE DATABASE for testing
CREATE DATABASE test_db;

-- Establishes a new connection
\connect test_db

-- CREATE SCHEMA \
CREATE SCHEMA private;

-- CREATE ROLE
CREATE ROLE admin WITH LOGIN PASSWORD 'P@ssw0rd';

-- GRANT ALL PRIVILEGES
GRANT ALL PRIVILEGES ON SCHEMA private TO admin;

-- CREATE TABLE | private.users
CREATE TABLE private.users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- CREATE TABLE | private.measurements
CREATE TABLE private.measurements (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  weight FLOAT NOT NULL,
  height FLOAT NOT NULL,
  body_fat FLOAT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES private.users(id)
);