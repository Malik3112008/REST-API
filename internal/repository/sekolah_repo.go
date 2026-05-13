package repository

import (
	"study/internal/model"
	"gorm.io/gorm"
)

type SekolahRepository interface {
	TambahGuru(guru *model.Guru) error
	AmbilSemuaGuru() ([]model.Guru, error)
	TambahKelas(kelas *model.Kelas) error
	AmbilSemuaKelas() ([]model.Kelas, error)
	TambahSiswa(siswa *model.Siswa) error
	AmbilSemuaSiswa() ([]model.Siswa, error)
}

type sekolahRepository struct {
	db *gorm.DB
}

func NewSekolahRepository(db *gorm.DB) SekolahRepository {
	return &sekolahRepository{db}
}

func (r *sekolahRepository) TambahGuru(guru *model.Guru) error {
	return r.db.Create(guru).Error
}

func (r *sekolahRepository) AmbilSemuaGuru() ([]model.Guru, error) {
	var daftarGuru []model.Guru
	err := r.db.Find(&daftarGuru).Error
	return daftarGuru, err
}

func (r *sekolahRepository) TambahKelas(kelas *model.Kelas) error {
	return r.db.Create(kelas).Error
}

func (r *sekolahRepository) AmbilSemuaKelas() ([]model.Kelas, error) {
	var daftarKelas []model.Kelas
	err := r.db.Preload("Guru").Preload("Siswa").Find(&daftarKelas).Error
	return daftarKelas, err
}

func (r *sekolahRepository) TambahSiswa(siswa *model.Siswa) error {
	return r.db.Create(siswa).Error
}

func (r *sekolahRepository) AmbilSemuaSiswa() ([]model.Siswa, error) {
	var daftarSiswa []model.Siswa
	err := r.db.Find(&daftarSiswa).Error
	return daftarSiswa, err
}
