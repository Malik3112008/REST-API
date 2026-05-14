package handler

import (
	"net/http"
	"strconv"
	"study/internal/model"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
)

type KelasHandler struct {
	uc usecase.SekolahUsecase
}

func NewKelasHandler(uc usecase.SekolahUsecase) *KelasHandler {
	return &KelasHandler{uc}
}

func (h *KelasHandler) BuatKelas(c *gin.Context) {
	var kelas model.Kelas
	if err := c.ShouldBindJSON(&kelas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	if err := h.uc.SimpanKelas(&kelas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data kelas: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, kelas)
}

func (h *KelasHandler) AmbilSemuaKelas(c *gin.Context) {
	daftarKelas, err := h.uc.DaftarKelas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar kelas: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarKelas)
}

func (h *KelasHandler) AmbilKelas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	kelas, err := h.uc.DetailKelas(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelas dengan ID tersebut tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, kelas)
}

func (h *KelasHandler) PerbaruiKelas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	var kelas model.Kelas
	if err := c.ShouldBindJSON(&kelas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	kelas.ID = uint(id)
	if err := h.uc.UbahKelas(&kelas); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "record not found" || err.Error() == "kelas tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": "Gagal memperbarui data kelas: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, kelas)
}

func (h *KelasHandler) HapusKelas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := h.uc.HapusKelas(uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "kelas tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kelas berhasil dihapus"})
}
