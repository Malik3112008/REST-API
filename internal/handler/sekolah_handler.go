package handler

import (
	"net/http"
	"strconv"
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

// --- PEMBERSIHAN ---
func (h *SekolahHandler) HapusSemuaData(c *gin.Context) {
	if err := h.uc.BersihkanSemuaData(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membersihkan data: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Semua data berhasil dihapus"})
}

// --- GURU ---
func (h *SekolahHandler) BuatGuru(c *gin.Context) {
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

func (h *SekolahHandler) AmbilSemuaGuru(c *gin.Context) {
	daftarGuru, err := h.uc.DaftarGuru()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar guru: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarGuru)
}

func (h *SekolahHandler) AmbilGuru(c *gin.Context) {
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

func (h *SekolahHandler) PerbaruiGuru(c *gin.Context) {
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

func (h *SekolahHandler) HapusGuru(c *gin.Context) {
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

// --- KELAS ---
func (h *SekolahHandler) BuatKelas(c *gin.Context) {
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

func (h *SekolahHandler) AmbilSemuaKelas(c *gin.Context) {
	daftarKelas, err := h.uc.DaftarKelas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar kelas: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarKelas)
}

func (h *SekolahHandler) AmbilKelas(c *gin.Context) {
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

func (h *SekolahHandler) PerbaruiKelas(c *gin.Context) {
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

func (h *SekolahHandler) HapusKelas(c *gin.Context) {
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

// --- SISWA ---
func (h *SekolahHandler) BuatSiswa(c *gin.Context) {
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

func (h *SekolahHandler) AmbilSemuaSiswa(c *gin.Context) {
	daftarSiswa, err := h.uc.DaftarSiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar siswa: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarSiswa)
}

func (h *SekolahHandler) AmbilSiswa(c *gin.Context) {
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

func (h *SekolahHandler) PerbaruiSiswa(c *gin.Context) {
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

func (h *SekolahHandler) HapusSiswa(c *gin.Context) {
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

// --- MATA PELAJARAN ---
func (h *SekolahHandler) BuatMapel(c *gin.Context) {
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

func (h *SekolahHandler) AmbilSemuaMapel(c *gin.Context) {
	daftarMapel, err := h.uc.DaftarMapel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar mata pelajaran: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, daftarMapel)
}

func (h *SekolahHandler) AmbilMapel(c *gin.Context) {
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

func (h *SekolahHandler) PerbaruiMapel(c *gin.Context) {
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

func (h *SekolahHandler) HapusMapel(c *gin.Context) {
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
