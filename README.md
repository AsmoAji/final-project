# 🏢 Sistem Pengelolaan Aset & Pemeliharaan Perusahaan

REST API berbasis Go (Gin + GORM + MySQL) untuk mencatat, melacak, dan merencanakan pemeliharaan aset perusahaan seperti mesin, peralatan kantor, dan lainnya. Sistem ini mendukung otentikasi JWT, role-based access control, serta laporan dan aktivitas.

## 📦 Teknologi

- **Go** (Golang)
- **Gin** (Web Framework)
- **GORM** (ORM)
- **MySQL** (Database)
- **JWT** (Authentication)
- Swagger/Postman (Dokumentasi API)

## 📁 Struktur Proyek

```
final-project/
├── controller/
├── service/
├── repository/
├── entity/
├── config/
├── middleware/
├── route.go
├── main.go
```

## 🔐 Fitur

### 🔧 User Management
- `POST /register` – Registrasi pengguna
- `POST /login` – Login dan mendapatkan token JWT
- Role: **admin** dan **user**

### 📦 Manajemen Aset (Items)
- `GET /items` – Lihat semua aset
- `POST /items` – Tambah aset baru *(admin only)*
- `PUT /items/:id` – Edit aset *(admin only)*
- `DELETE /items/:id` – Hapus aset *(admin only)*

### 📋 Aktivitas Pemeliharaan
- `GET /activities` – Lihat semua aktivitas
- `GET /activities/:id` – Lihat detail aktivitas
- `POST /activities` – Tambah aktivitas (misalnya: pemeliharaan aset)

## 🚀 Cara Menjalankan

### 1. Clone Repo
```bash
git clone https://github.com/<username>/final-project.git
cd final-project
```

### 2. Konfigurasi Database
Edit `config/db.go`:
```go
dsn := "root:<password>@tcp(127.0.0.1:3306)/asset_management_db?charset=utf8mb4&parseTime=True&loc=Local"
```

### 3. Jalankan
```bash
go mod tidy
go run main.go
```

## 📬 Contoh Request

### 🔐 Login
```http
POST /login
{
  "email": "admin@example.com",
  "password": "admin123"
}
```

**Response:**
```json
{ "token": "eyJhbGciOi..." }
```

### 📦 Tambah Item
```http
POST /items
Authorization: Bearer <token>

{
  "name": "Laptop Lenovo",
  "category": "Elektronik",
  "serial_number": "SN123456",
  "purchase_date": "2024-06-14",
  "location": "Gudang 1",
  "condition": "Baru"
}
```

## 📄 Lisensi

MIT License © 2025 - [Putranda Asmo Aji](https://github.com/<username>)