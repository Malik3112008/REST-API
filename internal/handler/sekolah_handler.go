package handler

import (
	"net/http"
	"study/internal/model"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SekolahHandler struct {
	uc usecase.SekolahUsecase
}

func NewSekolahHandler(uc usecase.SekolahUsecase) *SekolahHandler {
	return &SekolahHandler{uc}
}

func (h *SekolahHandler) BuatGuru(c *gin.Context) {
	var guru model.Guru
	if err := c.ShouldBindJSON(&guru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.SimpanGuru(&guru); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, guru)
}

func (h *SekolahHandler) AmbilGuru(c *gin.Context) {
	daftarGuru, err := h.uc.DaftarGuru()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, daftarGuru)
}

func (h *SekolahHandler) BuatKelas(c *gin.Context) {
	var kelas model.Kelas
	if err := c.ShouldBindJSON(&kelas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.SimpanKelas(&kelas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, kelas)
}

func (h *SekolahHandler) AmbilKelas(c *gin.Context) {
	daftarKelas, err := h.uc.DaftarKelas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, daftarKelas)
}

func (h *SekolahHandler) BuatSiswa(c *gin.Context) {
	var siswa model.Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.SimpanSiswa(&siswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, siswa)
}

func (h *SekolahHandler) AmbilSiswa(c *gin.Context) {
	daftarSiswa, err := h.uc.DaftarSiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, daftarSiswa)
}
