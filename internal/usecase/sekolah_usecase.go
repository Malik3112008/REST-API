package usecase

import (
	"study/internal/model"
	"study/internal/repository"
)

type SekolahUsecase interface {
	SimpanGuru(guru *model.Guru) error
	DaftarGuru() ([]model.Guru, error)
	SimpanKelas(kelas *model.Kelas) error
	DaftarKelas() ([]model.Kelas, error)
	SimpanSiswa(siswa *model.Siswa) error
	DaftarSiswa() ([]model.Siswa, error)
}

type sekolahUsecase struct {
	repo repository.SekolahRepository
}

func NewSekolahUsecase(repo repository.SekolahRepository) SekolahUsecase {
	return &sekolahUsecase{repo}
}

func (u *sekolahUsecase) SimpanGuru(guru *model.Guru) error {
	return u.repo.TambahGuru(guru)
}

func (u *sekolahUsecase) DaftarGuru() ([]model.Guru, error) {
	return u.repo.AmbilSemuaGuru()
}

func (u *sekolahUsecase) SimpanKelas(kelas *model.Kelas) error {
	return u.repo.TambahKelas(kelas)
}

func (u *sekolahUsecase) DaftarKelas() ([]model.Kelas, error) {
	return u.repo.AmbilSemuaKelas()
}

func (u *sekolahUsecase) SimpanSiswa(siswa *model.Siswa) error {
	return u.repo.TambahSiswa(siswa)
}

func (u *sekolahUsecase) DaftarSiswa() ([]model.Siswa, error) {
	return u.repo.AmbilSemuaSiswa()
}
