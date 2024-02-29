
-- foydalanuvchilar jadvali
CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    full_name VARCHAR(150),
    email VARCHAR(100),
    user_name VARCHAR(100) UNIQUE,
    user_password VARCHAR(100)
);

-- todo (vazifa) jadvali
CREATE TABLE todo (
    id VARCHAR(50) PRIMARY KEY,
    title VARCHAR(100),
    status VARCHAR(50),
    user_id VARCHAR(50) REFERENCES users(id)
);
