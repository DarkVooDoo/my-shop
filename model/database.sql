CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  firstname VARCHAR(40),
  lastname VARCHAR(40),
  phone VARCHAR,
  adresse VARCHAR,
  postal INT,
  email VARCHAR UNIQUE,
  access_token TEXT,
  salt INT DEFAULT FLOOR(RANDOM() * 9999) + 1000,
  ispro BOOL DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE category (
  id SERIAL PRIMARY KEY,
  name VARCHAR
);

CREATE TABLE etablishment (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(50),
  phone VARCHAR(13),
  adresse VARCHAR(150),
  postal INT,
  created_at date DEFAULT NOW(),
  user_id BIGINT REFERENCES users(users.id) ON DELETE CASCADE,
  category_id INT REFERENCES category(category.id)
);

/*CREATE TABLE message (
  id BIGSERIAL PRIMARY KEY,
  msg text,
  created_at date,
  etablishment_id bigint,
  from_id bigint,
  to_id bigint
);

CREATE TABLE workday (
  id BIGSERIAL PRIMARY KEY,
  from time,
  to time,
  weekday int,
  etablishment_id bigint
);

CREATE TABLE service (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100),
  timer time,
  price money,
  etablishment_id bigint
);

CREATE TABLE employee (
  id BIGSERIAL PRIMARY KEY,
  dayoff int[],
  etablishment_id bigint,
  user_id bigint
);

CREATE TABLE review (
  id BIGSERIAL,
  message text,
  star int,
  created_at date,
  etablishment_id bigint,
  user_id bigint,
  employee_id bigint
);

CREATE TABLE appointment_wa (
  id BIGSERIAL PRIMARY KEY,
  date TIMESTAMP,
  status VARCHAR,
  created_at TIMESTAMP,
  user_id integer,
  etablishment_id bigint,
  employee_id bigint
);

CREATE TABLE appointment (
  id BIGSERIAL PRIMARY KEY,
  "date" TIMESTAMP,
  firstname VARCHAR(40),
  lastname VARCHAR(40),
  email VARCHAR(100),
  status VARCHAR,
  create_at TIMESTAMP,
  etablishment_id bigint,
  employee_id bigint
);
*/
