-- Drop and Create User table
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(254) NOT NULL,
    username VARCHAR(50) NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX users_email_idx ON users (email);

-- Drop and Create Resume table
DROP TABLE IF EXISTS resumes;
CREATE TABLE resumes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    file_name VARCHAR(100) NOT NULL,
    is_master BOOLEAN NOT NULL DEFAULT FALSE,
    size INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);


