CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    release_year INT CHECK(release_year >= 1980 AND release_year <= 2024),
    price INT,
    total_page INT,
    thickness VARCHAR(50),
    category_id INT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
