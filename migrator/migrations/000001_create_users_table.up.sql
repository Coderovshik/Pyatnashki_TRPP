CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    user_password VARCHAR NOT NULL
);