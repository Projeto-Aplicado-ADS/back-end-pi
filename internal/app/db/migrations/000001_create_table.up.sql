CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(200) NOT NULL UNIQUE,
    phone VARCHAR(11) NOT NULL,
    password VARCHAR(255) NOT NULL, 
    is_admin BOOLEAN NOT NULL DEFAULT 0,
    birthday BIGINT NOT NULL,
    created_at BIGINT NOT NULL,
    update_at BIGINT NOT NULL DEFAULT 0
);