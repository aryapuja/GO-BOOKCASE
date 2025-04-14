CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255)
);