package config

import (
	"fmt"
	"log"
	"os"
	"study/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env tidak ditemukan, menggunakan variabel lingkungan")
	}

	// Mengambil kredensial database dari lingkungan
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL tidak diatur")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke database: %w", err)
	}

	// Auto Migrate semua skema
	log.Println("Menjalankan migrasi database...")
	err = db.AutoMigrate(
		&model.Guru{},
		&model.Kelas{},
		&model.Siswa{},
		&model.MataPelajaran{},
	)
	if err != nil {
		return nil, fmt.Errorf("gagal migrasi database: %w", err)
	}

	log.Println("Migrasi database selesai dengan sukses.")
	return db, nil
}
