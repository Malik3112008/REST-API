# Dokumentasi API Manajemen Sekolah

Dokumentasi ini menjelaskan secara rinci seluruh endpoint API yang tersedia untuk aplikasi Manajemen Sekolah. API ini dibangun dengan Go dan Gin framework, mendukung format JSON untuk *request* dan *response*.

Base URL (default lokal): `http://localhost:8080`

Aplikasi juga menyediakan OpenAPI (Swagger) viewer menggunakan Scalar yang dapat diakses melalui:
`GET /scalar`

---

## Tabel Konten
1. [Mata Pelajaran (Mapel)](#1-mata-pelajaran-mapel)
2. [Guru](#2-guru)
3. [Kelas](#3-kelas)
4. [Siswa](#4-siswa)

---

## 1. Mata Pelajaran (Mapel)

Mengelola data mata pelajaran yang ada di sekolah.

### 1.1 Buat Mata Pelajaran Baru
- **URL**: `/mapel`
- **Method**: `POST`
- **Request Body** (JSON):
  ```json
  {
    "nama": "Matematika"
  }
  ```
- **Response Success** (201 Created):
  ```json
  {
    "id": 1,
    "nama": "Matematika",
    "created_at": "2023-10-01T12:00:00Z",
    "updated_at": "2023-10-01T12:00:00Z"
  }
  ```
- **Response Error** (400 Bad Request / 500 Internal Server Error):
  ```json
  {
    "error": "Pesan error"
  }
  ```

### 1.2 Ambil Semua Mata Pelajaran
- **URL**: `/mapel`
- **Method**: `GET`
- **Response Success** (200 OK):
  ```json
  [
    {
      "id": 1,
      "nama": "Matematika",
      "created_at": "...",
      "updated_at": "..."
    }
  ]
  ```

### 1.3 Ambil Detail Mata Pelajaran
- **URL**: `/mapel/:id`
- **Method**: `GET`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK):
  ```json
  {
    "id": 1,
    "nama": "Matematika",
    "created_at": "...",
    "updated_at": "..."
  }
  ```
- **Response Error** (404 Not Found):
  ```json
  {
    "error": "Mata Pelajaran dengan ID tersebut tidak ditemukan"
  }
  ```

### 1.4 Perbarui Data Mata Pelajaran
- **URL**: `/mapel/:id`
- **Method**: `PUT`
- **URL Params**: `id`=[integer]
- **Request Body** (JSON):
  ```json
  {
    "nama": "Matematika Lanjut"
  }
  ```
- **Response Success** (200 OK):
  ```json
  {
    "id": 1,
    "nama": "Matematika Lanjut",
    "created_at": "...",
    "updated_at": "..."
  }
  ```

### 1.5 Hapus Mata Pelajaran
- **URL**: `/mapel/:id`
- **Method**: `DELETE`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK):
  ```json
  {
    "message": "Mata Pelajaran berhasil dihapus"
  }
  ```

---

## 2. Guru

Mengelola data guru, yang memiliki relasi ke Mata Pelajaran.

### 2.1 Buat Guru Baru
- **URL**: `/guru`
- **Method**: `POST`
- **Request Body** (JSON):
  ```json
  {
    "nama": "Budi Santoso",
    "email": "budi.santoso@sekolah.com",
    "mata_pelajaran_id": 1
  }
  ```
- **Response Success** (201 Created):
  ```json
  {
    "id": 1,
    "nama": "Budi Santoso",
    "email": "budi.santoso@sekolah.com",
    "mata_pelajaran_id": 1,
    "mata_pelajaran": { ... },
    "created_at": "...",
    "updated_at": "..."
  }
  ```

### 2.2 Ambil Semua Guru
- **URL**: `/guru`
- **Method**: `GET`
- **Response Success** (200 OK): Array of Object Guru.

### 2.3 Ambil Detail Guru
- **URL**: `/guru/:id`
- **Method**: `GET`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK): Object Guru.
- **Response Error** (404 Not Found):
  ```json
  {
    "error": "Guru dengan ID tersebut tidak ditemukan"
  }
  ```

### 2.4 Perbarui Data Guru
- **URL**: `/guru/:id`
- **Method**: `PUT`
- **URL Params**: `id`=[integer]
- **Request Body** (JSON):
  ```json
  {
    "nama": "Budi Santoso S.Pd",
    "email": "budi.santoso@sekolah.com",
    "mata_pelajaran_id": 2
  }
  ```
- **Response Success** (200 OK): Object Guru yang diperbarui.

### 2.5 Hapus Guru
- **URL**: `/guru/:id`
- **Method**: `DELETE`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK):
  ```json
  {
    "message": "Guru berhasil dihapus"
  }
  ```

---

## 3. Kelas

Mengelola data kelas, yang terkait dengan Guru (wali kelas).

### 3.1 Buat Kelas Baru
- **URL**: `/kelas`
- **Method**: `POST`
- **Request Body** (JSON):
  ```json
  {
    "nama": "10-A",
    "guru_id": 1
  }
  ```
- **Response Success** (201 Created):
  ```json
  {
    "id": 1,
    "nama": "10-A",
    "guru_id": 1,
    "guru": { ... },
    "siswa": [],
    "created_at": "...",
    "updated_at": "..."
  }
  ```

### 3.2 Ambil Semua Kelas
- **URL**: `/kelas`
- **Method**: `GET`
- **Response Success** (200 OK): Array of Object Kelas (termasuk list siswa jika preloaded).

### 3.3 Ambil Detail Kelas
- **URL**: `/kelas/:id`
- **Method**: `GET`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK): Object Kelas.

### 3.4 Perbarui Data Kelas
- **URL**: `/kelas/:id`
- **Method**: `PUT`
- **URL Params**: `id`=[integer]
- **Request Body** (JSON):
  ```json
  {
    "nama": "10-IPA-1",
    "guru_id": 2
  }
  ```
- **Response Success** (200 OK): Object Kelas yang diperbarui.

### 3.5 Hapus Kelas
- **URL**: `/kelas/:id`
- **Method**: `DELETE`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK):
  ```json
  {
    "message": "Kelas berhasil dihapus"
  }
  ```

---

## 4. Siswa

Mengelola data siswa, yang terkait dengan Kelas tempat siswa tersebut berada.

### 4.1 Buat Siswa Baru
- **URL**: `/siswa`
- **Method**: `POST`
- **Request Body** (JSON):
  ```json
  {
    "nama": "Ahmad Dani",
    "email": "ahmad.dani@siswa.com",
    "kelas_id": 1
  }
  ```
- **Response Success** (201 Created):
  ```json
  {
    "id": 1,
    "nama": "Ahmad Dani",
    "email": "ahmad.dani@siswa.com",
    "kelas_id": 1,
    "kelas": { ... },
    "created_at": "...",
    "updated_at": "..."
  }
  ```

### 4.2 Ambil Semua Siswa
- **URL**: `/siswa`
- **Method**: `GET`
- **Response Success** (200 OK): Array of Object Siswa.

### 4.3 Ambil Detail Siswa
- **URL**: `/siswa/:id`
- **Method**: `GET`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK): Object Siswa.

### 4.4 Perbarui Data Siswa
- **URL**: `/siswa/:id`
- **Method**: `PUT`
- **URL Params**: `id`=[integer]
- **Request Body** (JSON):
  ```json
  {
    "nama": "Ahmad Dani Setiawan",
    "email": "ahmad.dani@siswa.com",
    "kelas_id": 2
  }
  ```
- **Response Success** (200 OK): Object Siswa yang diperbarui.

### 4.5 Hapus Siswa
- **URL**: `/siswa/:id`
- **Method**: `DELETE`
- **URL Params**: `id`=[integer]
- **Response Success** (200 OK):
  ```json
  {
    "message": "Siswa berhasil dihapus"
  }
  ```

---

## Handling Error Secara Umum
API menggunakan HTTP status code standar untuk menyatakan hasil dari suatu *request*:
- **200 OK**: Request berhasil.
- **201 Created**: Data berhasil dibuat.
- **400 Bad Request**: Request tidak valid (misal: JSON body salah, parameter salah).
- **404 Not Found**: Data atau *endpoint* tidak ditemukan.
- **500 Internal Server Error**: Terjadi kesalahan pada server (misal: database tidak terkoneksi, constraint error).

Format error response:
```json
{
  "error": "Deskripsi error"
}
```
