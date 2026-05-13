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

	// --- ENDPOINTS SISTEM ---
	r.DELETE("/sistem/reset", h.HapusSemuaData)

	// --- ENDPOINTS GURU ---
	r.POST("/guru", h.BuatGuru)
	r.GET("/guru", h.AmbilSemuaGuru)
	r.GET("/guru/:id", h.AmbilGuru)
	r.PUT("/guru/:id", h.PerbaruiGuru)
	r.DELETE("/guru/:id", h.HapusGuru)

	// --- ENDPOINTS KELAS ---
	r.POST("/kelas", h.BuatKelas)
	r.GET("/kelas", h.AmbilSemuaKelas)
	r.GET("/kelas/:id", h.AmbilKelas)
	r.PUT("/kelas/:id", h.PerbaruiKelas)
	r.DELETE("/kelas/:id", h.HapusKelas)

	// --- ENDPOINTS SISWA ---
	r.POST("/siswa", h.BuatSiswa)
	r.GET("/siswa", h.AmbilSemuaSiswa)
	r.GET("/siswa/:id", h.AmbilSiswa)
	r.PUT("/siswa/:id", h.PerbaruiSiswa)
	r.DELETE("/siswa/:id", h.HapusSiswa)

	// --- ENDPOINTS MATA PELAJARAN ---
	r.POST("/mapel", h.BuatMapel)
	r.GET("/mapel", h.AmbilSemuaMapel)
	r.GET("/mapel/:id", h.AmbilMapel)
	r.PUT("/mapel/:id", h.PerbaruiMapel)
	r.DELETE("/mapel/:id", h.HapusMapel)

	log.Println("Server API Manajemen Sekolah berjalan di :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
