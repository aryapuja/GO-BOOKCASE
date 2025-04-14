\c go_bookcase;

CREATE TABLE IF NOT EXISTS books (
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

-- Insert Buku untuk Kategori 'Fiction' (ID = 1)
INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, created_by)
VALUES
    ('Laskar Pelangi', 'Novel tentang perjuangan sekelompok anak-anak di Belitung', '-', 2005, 100000, 400, 1, 'admin'),
    ('Bumi Manusia', 'Novel yang menceritakan sejarah Indonesia melalui kisah cinta', '-', 1980, 120000, 500, 1, 'admin'),
    ('Dilan: Dia adalah Dilanku', 'Kisah cinta remaja di Bandung tahun 1990an', '-', 2014, 90000, 350, 1, 'admin');

-- Insert Buku untuk Kategori 'Non-Fiction' (ID = 2)
INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, created_by)
VALUES
    ('Sapiens: A Brief History of Humankind', 'Buku tentang sejarah manusia dan evolusinya', '-', 2015, 150000, 600, 2, 'admin'),
    ('Catatan Seorang Demonstran', 'Buku tentang pengalaman seorang aktivis mahasiswa dalam gerakan reformasi 1998', '-', 1998, 90000, 270, 2, 'admin'),
    ('Mereka yang Terlupakan', 'Menceritakan kehidupan sosial di Indonesia pada masa lalu', '-', 2012, 95000, 330, 2, 'admin');

-- Insert Buku untuk Kategori 'Science' (ID = 3)
INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, created_by)
VALUES
    ('Ilmu Pengetahuan Alam untuk Anak', 'Buku sains dasar untuk anak-anak Indonesia', '-', 2020, 80000, 120, 3, 'admin'),
    ('Teknologi Masa Depan', 'Buku yang membahas perkembangan teknologi dan dampaknya', '-', 2021, 140000, 400, 3, 'admin'),
    ('Kehidupan di Bawah Laut', 'Buku tentang ekosistem laut Indonesia', '-', 2018, 95000, 300, 3, 'admin');

-- Insert Buku untuk Kategori 'History' (ID = 4)
INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, created_by)
VALUES
    ('Sejarah Indonesia: Dari Masa ke Masa', 'Buku yang menceritakan perjalanan sejarah Indonesia', '-', 2016, 110000, 450, 4, 'admin'),
    ('Perjuangan Pahlawan Nasional', 'Buku tentang pahlawan-pahlawan Indonesia yang berjuang untuk kemerdekaan', '-', 2000, 95000, 500, 4, 'admin'),
    ('Kehidupan Zaman Kolonial', 'Buku tentang kehidupan masyarakat Indonesia pada masa kolonial', '-', 2010, 100000, 380, 4, 'admin');

-- Insert Buku untuk Kategori 'Biography' (ID = 5)
INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, created_by)
VALUES
    ('Bung Karno: Sang Fajar', 'Biografi Presiden pertama Indonesia, Sukarno', '-', 2011, 140000, 450, 5, 'admin'),
    ('Gaya Hidup Haji Agus Salim', 'Biografi tokoh kemerdekaan Indonesia, Agus Salim', '-', 2005, 90000, 250, 5, 'admin'),
    ('Ahok: Dari Penjara ke Istana', 'Biografi mantan gubernur DKI Jakarta, Ahok', '-', 2018, 130000, 400, 5, 'admin');

-- Insert Buku untuk Kategori 'Technology' (ID = 6)
INSERT INTO books (title, description, image_url, release_year, price, total_page, category_id, created_by)
VALUES
    ('Revolusi Industri 4.0', 'Buku yang menjelaskan perkembangan revolusi industri saat ini', '-', 2017, 120000, 350, 6, 'admin'),
    ('Teknologi 5G di Indonesia', 'Buku tentang pengaruh teknologi 5G di Indonesia', '-', 2020, 100000, 280, 6, 'admin'),
    ('AI untuk Pemula', 'Buku tentang pengenalan teknologi Artificial Intelligence', '-', 2021, 110000, 300, 6, 'admin');
