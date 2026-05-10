# API Manajemen Proyek

REST API siap produksi yang dibangun dengan Go, Gin, dan GORM.

## Arsitektur
Proyek ini mengikuti **Arsitektur Berlapis (Clean Architecture)**:
- **Handler**: Menangani request dan response HTTP.
- **Usecase/Service**: Berisi logika bisnis.
- **Repository**: Menangani operasi database.
- **Model**: Mendefinisikan struktur data dan skema database.

## Stack Teknologi
- **Bahasa**: Go 1.21+
- **Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: SQLite (Migrasi Otomatis)

## Memulai

### Prasyarat
- Go sudah terinstal (1.21 atau lebih tinggi)

### Menjalankan Aplikasi
1. Instal dependensi:
   ```bash
   go mod tidy
   ```
2. Jalankan server:
   ```bash
   go run cmd/api/main.go
   ```
   Server akan berjalan di `http://localhost:8080`.

## Endpoint API

### Tugas (Tasks)
- **Buat Tugas**
  - `POST /tasks`
  - Body:
    ```json
    {
      "title": "Perbaiki bug",
      "project_id": 1,
      "assignee_id": 1
    }
    ```
- **Ambil Tugas berdasarkan ID**
  - `GET /tasks/:id`
- **Perbarui Status Tugas**
  - `PATCH /tasks/:id/status`
  - Body:
    ```json
    {
      "status": "selesai"
    }
    ```

## Pengujian
Untuk menguji API, Anda dapat menggunakan `curl` atau klien API apa pun seperti Postman.

Contoh `curl` untuk membuat tugas:
```bash
curl -X POST http://localhost:8080/tasks \
     -H "Content-Type: application/json" \
     -d '{"title": "Tugas Awal", "project_id": 1, "assignee_id": 1}'
```
