package handler

import (
	"net/http"
	"strconv"
	"study/internal/model"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SiswaHandler struct {
	uc usecase.SekolahUsecase
}

func NewSiswaHandler(uc usecase.SekolahUsecase) *SiswaHandler {
	return &SiswaHandler{uc}
}

func (h *SiswaHandler) BuatSiswa(c *gin.Context) {
	var siswa model.Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	if err := h.uc.SimpanSiswa(&siswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data siswa: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, siswa)
}

func (h *SiswaHandler) AmbilSemuaSiswa(c *gin.Context) {
	daftarSiswa, err := h.uc.DaftarSiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar siswa: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarSiswa)
}

func (h *SiswaHandler) AmbilSiswa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	siswa, err := h.uc.DetailSiswa(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Siswa dengan ID tersebut tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, siswa)
}

func (h *SiswaHandler) PerbaruiSiswa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	var siswa model.Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	siswa.ID = uint(id)
	if err := h.uc.UbahSiswa(&siswa); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "record not found" || err.Error() == "siswa tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": "Gagal memperbarui data siswa: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, siswa)
}

func (h *SiswaHandler) HapusSiswa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := h.uc.HapusSiswa(uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "siswa tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Siswa berhasil dihapus"})
}
