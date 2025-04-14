# GO-BOOKCASE (Repository Kuis SanberCode)

Go Bookcase adalah aplikasi API untuk mengelola informasi buku, kategori, dan pengguna dengan autentikasi menggunakan **JWT**

Deploy: http://go-bookcase-production.up.railway.app (Bug route login + swagger tidak terbaca, tapi route dengan validasi JWT sudah aktif)

## Kegunaan

- **Manajemen Buku**: Menambahkan, mengupdate, menghapus, dan menampilkan buku-buku yang ada.
- **Manajemen Kategori**: Menambah, mengupdate, menghapus, dan menampilkan kategori buku.
- **Autentikasi Pengguna**: Menggunakan **JWT (JSON Web Token)** untuk autentikasi dan otorisasi akses API.

## Persyaratan

- **Go (Golang)** version 1.x atau lebih tinggi.
- **PostgreSQL** sebagai database.
- **Gin Gonic** sebagai web framework.
- **JWT** untuk autentikasi.
- **Swagger** version 1.0.0

## Instalasi

1. Clone repositori ini:

    ```bash
    git clone https://github.com/username/go-bookcase.git
    cd go-bookcase
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Siapkan file `.env` untuk variabel lingkungan:

    ```env
    DATABASE_URL=postgres://username:password@localhost:5432/go_bookcase
    SECRET_KEY=your_secret_key
    PORT=8080
    ```

4. Jalankan aplikasi:

    ```bash
    go run main.go
    ```

    Aplikasi akan berjalan di `http://localhost:8080`.

## API Endpoints

### **1. /api/books**

- **GET**: Mengambil semua buku.
- **POST**: Menambahkan buku baru.
  
### **2. /api/books/{id}**

- **GET**: Mengambil detail buku berdasarkan ID.
- **PUT**: Mengupdate buku berdasarkan ID.
- **DELETE**: Menghapus buku berdasarkan ID.

### **3. /api/categories**

- **GET**: Mengambil semua kategori.
- **POST**: Menambahkan kategori baru.

### **4. /api/categories/{id}**

- **GET**: Mengambil detail kategori berdasarkan ID.
- **PUT**: Mengupdate kategori berdasarkan ID.
- **DELETE**: Menghapus kategori berdasarkan ID.

### **5. /api/categories/{id}/books**

- **GET**: Mengambil list buku berdasarkan ID kategori.

### **6. /api/users/login**

- **POST**: Login dan dapatkan token JWT.

## Autentikasi JWT

Untuk mengakses API yang memerlukan autentikasi, kirimkan **token JWT** di header `Authorization` dengan format **Bearer token**.

### **Contoh Request untuk Login**:

**POST** `/api/users/login`

```json
{
  "username": "admin",
  "password": "password"
}
