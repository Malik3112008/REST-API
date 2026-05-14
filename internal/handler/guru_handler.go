package handler

import (
	"net/http"
	"strconv"
	"study/internal/model"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
)

type GuruHandler struct {
	uc usecase.SekolahUsecase
}

func NewGuruHandler(uc usecase.SekolahUsecase) *GuruHandler {
	return &GuruHandler{uc}
}

func (h *GuruHandler) BuatGuru(c *gin.Context) {
	var guru model.Guru
	if err := c.ShouldBindJSON(&guru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	if err := h.uc.SimpanGuru(&guru); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data guru: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, guru)
}

func (h *GuruHandler) AmbilSemuaGuru(c *gin.Context) {
	daftarGuru, err := h.uc.DaftarGuru()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar guru: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarGuru)
}

func (h *GuruHandler) AmbilGuru(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	guru, err := h.uc.DetailGuru(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guru dengan ID tersebut tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, guru)
}

func (h *GuruHandler) PerbaruiGuru(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	var guru model.Guru
	if err := c.ShouldBindJSON(&guru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	guru.ID = uint(id)
	if err := h.uc.UbahGuru(&guru); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "record not found" || err.Error() == "guru tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": "Gagal memperbarui data guru: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, guru)
}

func (h *GuruHandler) HapusGuru(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := h.uc.HapusGuru(uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "guru tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Guru berhasil dihapus"})
}
