# Panduan Demo API Manajemen Sekolah

Dokumentasi ini menjelaskan cara mendemonstrasikan seluruh fitur API dari hulu ke hilir, mencakup seluruh entitas yang ada (Mata Pelajaran, Guru, Kelas, dan Siswa).

## Persiapan
1. Jalankan database PostgreSQL Anda.
2. Pastikan file `.env` sudah terisi dengan `DATABASE_URL` yang benar.
3. Jalankan server:
   ```bash
   go run cmd/api/main.go
   ```
4. Server akan berjalan di `http://localhost:8080`.

---

## Urutan Demo (Aliran Data)

Untuk mendemonstrasikan relasi antar data, gunakan urutan pemanggilan berikut:

### 1. Manajemen Mata Pelajaran (Mapel)
Langkah awal karena Guru membutuhkan Mata Pelajaran.
- **Create Mapel**: `POST /mapel`
  ```json
  { "nama": "Bahasa Indonesia" }
  ```
- **Get All**: `GET /mapel`

### 2. Manajemen Guru
Guru membutuhkan `mata_pelajaran_id`.
- **Create Guru**: `POST /guru`
  ```json
  {
    "nama": "Pak Slamet",
    "email": "slamet@sekolah.id",
    "mata_pelajaran_id": 1
  }
  ```
- **Get Detail**: `GET /guru/1` (Akan menampilkan data guru beserta detail mata pelajarannya).

### 3. Manajemen Kelas
Kelas membutuhkan `guru_id` sebagai wali kelas.
- **Create Kelas**: `POST /kelas`
  ```json
  {
    "nama": "X - IPA 1",
    "guru_id": 1
  }
  ```

### 4. Manajemen Siswa
Siswa membutuhkan `kelas_id`.
- **Create Siswa**: `POST /siswa`
  ```json
  {
    "nama": "Budi Santoso",
    "email": "budi@student.id",
    "kelas_id": 1
  }
  ```

---

## Fitur Unggulan untuk Didemokan

### 1. Relasi Otomatis (Eager Loading)
Coba panggil `GET /kelas/1`. 
**Hasil yang diharapkan**: API tidak hanya memberikan nama kelas, tapi secara otomatis melampirkan:
- Data **Guru** (Wali Kelas).
- Daftar seluruh **Siswa** yang terdaftar di kelas tersebut.

### 2. Validasi & Error Handling
- Coba masukkan **Email yang sama** dua kali pada Siswa atau Guru. 
**Hasil**: API akan mengembalikan error `500` atau `400` karena adanya constraint `UNIQUE` di database.
- Coba akses ID yang tidak ada (misal: `GET /siswa/999`).
**Hasil**: API mengembalikan pesan `Siswa tidak ditemukan`.

### 3. Dokumentasi Interaktif (Scalar)
Buka browser dan akses:
`http://localhost:8080/scalar`
Ini akan menampilkan dokumentasi visual yang bisa digunakan untuk mencoba API secara langsung tanpa Postman.

---

## Pembersihan (Cleanup)
Karena fitur "Reset Sistem" telah dihapus demi keamanan, penghapusan data harus dilakukan secara manual per entitas:
- `DELETE /siswa/{id}`
- `DELETE /kelas/{id}`
- `DELETE /guru/{id}`
- `DELETE /mapel/{id}`
