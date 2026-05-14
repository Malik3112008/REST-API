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
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env tidak ditemukan, menggunakan variabel lingkungan")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL tidak diatur")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke database: %w", err)
	}

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
