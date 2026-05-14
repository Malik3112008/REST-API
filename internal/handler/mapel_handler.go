package handler

import (
	"net/http"
	"strconv"
	"study/internal/model"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MapelHandler struct {
	uc usecase.SekolahUsecase
}

func NewMapelHandler(uc usecase.SekolahUsecase) *MapelHandler {
	return &MapelHandler{uc}
}

func (h *MapelHandler) BuatMapel(c *gin.Context) {
	var mapel model.MataPelajaran
	if err := c.ShouldBindJSON(&mapel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	if err := h.uc.SimpanMapel(&mapel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data mata pelajaran: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mapel)
}

func (h *MapelHandler) AmbilSemuaMapel(c *gin.Context) {
	daftarMapel, err := h.uc.DaftarMapel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar mata pelajaran: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarMapel)
}

func (h *MapelHandler) AmbilMapel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	mapel, err := h.uc.DetailMapel(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mata Pelajaran dengan ID tersebut tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, mapel)
}

func (h *MapelHandler) PerbaruiMapel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	var mapel model.MataPelajaran
	if err := c.ShouldBindJSON(&mapel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	mapel.ID = uint(id)
	if err := h.uc.UbahMapel(&mapel); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "record not found" || err.Error() == "mata pelajaran tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": "Gagal memperbarui data mata pelajaran: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapel)
}

func (h *MapelHandler) HapusMapel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := h.uc.HapusMapel(uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "mata pelajaran tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mata Pelajaran berhasil dihapus"})
}
