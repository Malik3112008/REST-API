# Dokumentasi Proyek Manajemen Sekolah

Proyek ini adalah API Manajemen Sekolah yang dibangun menggunakan bahasa pemrograman Go dengan framework Gin dan GORM. Proyek ini mengikuti arsitektur **Clean Architecture** (atau Layered Architecture) untuk memisahkan tanggung jawab antar komponen.

## Struktur Direktori

### 1. `cmd/api/`
- **`main.go`**: Entry point aplikasi. Berisi konfigurasi server, inisialisasi database, dependency injection, dan pendaftaran route API.

### 2. `internal/handler/`
Berisi logika untuk menangani request HTTP, validasi input awal, dan mengembalikan response (JSON).
- **`sekolah_handler.go`**: Berisi definisi struct utama `SekolahHandler` dan constructor-nya.
- **`siswa_handler.go`**: Menangani endpoint yang berkaitan dengan data **Siswa** (Create, Read, Update, Delete).
- **`guru_handler.go`**: Menangani endpoint yang berkaitan dengan data **Guru**.
- **`kelas_handler.go`**: Menangani endpoint yang berkaitan dengan data **Kelas**.
- **`mapel_handler.go`**: Menangani endpoint yang berkaitan dengan data **Mata Pelajaran**.

### 3. `internal/usecase/`
Berisi **Business Logic** aplikasi. Lapisan ini menjadi jembatan antara Handler dan Repository.
- **`sekolah_usecase.go`**: Mendefinisikan aturan bisnis (misalnya: validasi tambahan sebelum simpan ke database).

### 4. `internal/repository/`
Berisi logika untuk akses ke data (Database).
- **`sekolah_repo.go`**: Menggunakan GORM untuk melakukan operasi CRUD ke database PostgreSQL/MySQL.

### 5. `internal/model/`
- **`models.go`**: Definisi struct atau skema data yang digunakan di seluruh aplikasi dan dipetakan ke tabel database.

### 6. `config/`
- **`database.go`**: Konfigurasi koneksi ke database.

### 7. `pkg/`
Berisi library atau utility yang bersifat umum dan bisa digunakan kembali di project lain (Helper).

## Arsitektur Data Flow
1. **Request** masuk melalui `main.go` (Router).
2. Diteruskan ke **Handler** (Misal: `siswa_handler.go`) untuk parsing data JSON.
3. Handler memanggil fungsi di **Usecase** untuk memproses logika bisnis.
4. Usecase memanggil **Repository** untuk menyimpan atau mengambil data dari database.
5. **Response** dikirimkan kembali dalam format JSON.

## Cara Menjalankan
1. Pastikan file `.env` sudah dikonfigurasi.
2. Jalankan perintah:
   ```bash
   go run cmd/api/main.go
   ```
