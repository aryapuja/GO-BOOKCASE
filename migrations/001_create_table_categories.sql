-- Membuat database jika belum ada
-- CREATE DATABASE IF NOT EXISTS go_bookcase;

-- Gunakan database go_bookcase
-- \c go_bookcase;

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255)
);

INSERT INTO categories (name, created_by)
VALUES
    ('Fiction', 'admin'),
    ('Non-Fiction', 'admin'),
    ('Science', 'admin'),
    ('History', 'admin'),
    ('Biography', 'admin'),
    ('Technology', 'admin');