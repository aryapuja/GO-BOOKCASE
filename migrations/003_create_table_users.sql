-- \c go_bookcase;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255)
);

INSERT INTO users (username, password, created_by)
VALUES
    ('admin', '$2a$12$Jb6pkn2oTf0O.vOReZChDii1FlwZ8Wz5v7rzY6eOP44cmYzjexOZC', 'admin');