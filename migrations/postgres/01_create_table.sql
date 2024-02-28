-- statuslar jadvali
CREATE TABLE status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE
);

-- foydalanuvchilar jadvali
CREATE TABLE users (
    id VARCHAR(34) PRIMARY KEY,
    full_name VARCHAR(150),
    email VARCHAR(100),
    user_name VARCHAR(100) UNIQUE,
    user_password VARCHAR(100)
);

-- todo (vazifa) jadvali
CREATE TABLE todo (
    id VARCHAR(34) PRIMARY KEY,
    title VARCHAR(100),
    status_id INTEGER REFERENCES status(id),
    user_id VARCHAR(34) REFERENCES users(id)
);
