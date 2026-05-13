package main

import (
	"log"
	"study/config"
	"study/internal/handler"
	"study/internal/repository"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/watchakorn-18k/scalar-go"
	scalargin "github.com/watchakorn-18k/scalar-go/middleware/gin"
)

func main() {
	// Inisialisasi Database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	// Dependency Injection
	repo := repository.NewSekolahRepository(db)
	uc := usecase.NewSekolahUsecase(repo)
	h := handler.NewSekolahHandler(uc)

	// Setup Router
	r := gin.Default()

	// Dokumentasi Scalar
	r.GET("/scalar", scalargin.Handler(&scalar.Options{
		SpecURL: "./openapi.yaml",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Dokumentasi API Manajemen Sekolah",
		},
		DarkMode: true,
	}))

	// Endpoints
	r.POST("/guru", h.BuatGuru)
	r.GET("/guru", h.AmbilGuru)
	
	r.POST("/kelas", h.BuatKelas)
	r.GET("/kelas", h.AmbilKelas)
	
	r.POST("/siswa", h.BuatSiswa)
	r.GET("/siswa", h.AmbilSiswa)

	log.Println("Server API Manajemen Sekolah berjalan di :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
