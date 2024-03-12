-- users jadvali
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name VARCHAR(255) NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    user_password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
);

-- tasks jadvali
CREATE TABLE tasks (
    task_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    status VARCHAR(50),
    user_id UUID REFERENCES users(id)
);

-- comments jadvali
CREATE TABLE comments (
    comment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id UUID REFERENCES tasks(task_id),
    user_id UUID REFERENCES users(id),
    text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- user_roles jadvali
CREATE TABLE user_roles (
    role_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_name VARCHAR(100) NOT NULL
);

-- user_roles_users bog'lovchi jadvali
CREATE TABLE user_roles_users (
    user_id UUID REFERENCES users(id),
    role_id UUID REFERENCES user_roles(role_id),
    PRIMARY KEY (user_id, role_id)
);
