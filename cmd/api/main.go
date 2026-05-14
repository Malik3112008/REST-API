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
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	repo := repository.NewSekolahRepository(db)
	uc := usecase.NewSekolahUsecase(repo)
	
	guruH := handler.NewGuruHandler(uc)
	kelasH := handler.NewKelasHandler(uc)
	siswaH := handler.NewSiswaHandler(uc)
	mapelH := handler.NewMapelHandler(uc)

	r := gin.Default()

	r.GET("/scalar", scalargin.Handler(&scalar.Options{
		SpecURL: "./openapi.yaml",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Dokumentasi API Manajemen Sekolah",
		},
		DarkMode: true,
	}))

	r.POST("/guru", guruH.BuatGuru)
	r.GET("/guru", guruH.AmbilSemuaGuru)
	r.GET("/guru/:id", guruH.AmbilGuru)
	r.PUT("/guru/:id", guruH.PerbaruiGuru)
	r.DELETE("/guru/:id", guruH.HapusGuru)

	r.POST("/kelas", kelasH.BuatKelas)
	r.GET("/kelas", kelasH.AmbilSemuaKelas)
	r.GET("/kelas/:id", kelasH.AmbilKelas)
	r.PUT("/kelas/:id", kelasH.PerbaruiKelas)
	r.DELETE("/kelas/:id", kelasH.HapusKelas)

	r.POST("/siswa", siswaH.BuatSiswa)
	r.GET("/siswa", siswaH.AmbilSemuaSiswa)
	r.GET("/siswa/:id", siswaH.AmbilSiswa)
	r.PUT("/siswa/:id", siswaH.PerbaruiSiswa)
	r.DELETE("/siswa/:id", siswaH.HapusSiswa)

	r.POST("/mapel", mapelH.BuatMapel)
	r.GET("/mapel", mapelH.AmbilSemuaMapel)
	r.GET("/mapel/:id", mapelH.AmbilMapel)
	r.PUT("/mapel/:id", mapelH.PerbaruiMapel)
	r.DELETE("/mapel/:id", mapelH.HapusMapel)

	log.Println("Server API Manajemen Sekolah berjalan di :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
