package usecase

import (
	"study/internal/model"
	"study/internal/repository"
)

type SekolahUsecase interface {
	// Guru
	SimpanGuru(guru *model.Guru) error
	DaftarGuru() ([]model.Guru, error)
	DetailGuru(id uint) (model.Guru, error)
	UbahGuru(guru *model.Guru) error
	HapusGuru(id uint) error

	// Kelas
	SimpanKelas(kelas *model.Kelas) error
	DaftarKelas() ([]model.Kelas, error)
	DetailKelas(id uint) (model.Kelas, error)
	UbahKelas(kelas *model.Kelas) error
	HapusKelas(id uint) error

	// Siswa
	SimpanSiswa(siswa *model.Siswa) error
	DaftarSiswa() ([]model.Siswa, error)
	DetailSiswa(id uint) (model.Siswa, error)
	UbahSiswa(siswa *model.Siswa) error
	HapusSiswa(id uint) error

	// Mata Pelajaran
	SimpanMapel(mapel *model.MataPelajaran) error
	DaftarMapel() ([]model.MataPelajaran, error)
	DetailMapel(id uint) (model.MataPelajaran, error)
	UbahMapel(mapel *model.MataPelajaran) error
	HapusMapel(id uint) error

	// Pembersihan
	BersihkanSemuaData() error
}

type sekolahUsecase struct {
	repo repository.SekolahRepository
}

func NewSekolahUsecase(repo repository.SekolahRepository) SekolahUsecase {
	return &sekolahUsecase{repo}
}

// --- GURU ---
func (u *sekolahUsecase) SimpanGuru(guru *model.Guru) error {
	return u.repo.TambahGuru(guru)
}

func (u *sekolahUsecase) DaftarGuru() ([]model.Guru, error) {
	return u.repo.AmbilSemuaGuru()
}

func (u *sekolahUsecase) DetailGuru(id uint) (model.Guru, error) {
	return u.repo.AmbilGuruBerdasarkanID(id)
}

func (u *sekolahUsecase) UbahGuru(guru *model.Guru) error {
	return u.repo.PerbaruiGuru(guru)
}

func (u *sekolahUsecase) HapusGuru(id uint) error {
	return u.repo.HapusGuru(id)
}

// --- KELAS ---
func (u *sekolahUsecase) SimpanKelas(kelas *model.Kelas) error {
	return u.repo.TambahKelas(kelas)
}

func (u *sekolahUsecase) DaftarKelas() ([]model.Kelas, error) {
	return u.repo.AmbilSemuaKelas()
}

func (u *sekolahUsecase) DetailKelas(id uint) (model.Kelas, error) {
	return u.repo.AmbilKelasBerdasarkanID(id)
}

func (u *sekolahUsecase) UbahKelas(kelas *model.Kelas) error {
	return u.repo.PerbaruiKelas(kelas)
}

func (u *sekolahUsecase) HapusKelas(id uint) error {
	return u.repo.HapusKelas(id)
}

// --- SISWA ---
func (u *sekolahUsecase) SimpanSiswa(siswa *model.Siswa) error {
	return u.repo.TambahSiswa(siswa)
}

func (u *sekolahUsecase) DaftarSiswa() ([]model.Siswa, error) {
	return u.repo.AmbilSemuaSiswa()
}

func (u *sekolahUsecase) DetailSiswa(id uint) (model.Siswa, error) {
	return u.repo.AmbilSiswaBerdasarkanID(id)
}

func (u *sekolahUsecase) UbahSiswa(siswa *model.Siswa) error {
	return u.repo.PerbaruiSiswa(siswa)
}

func (u *sekolahUsecase) HapusSiswa(id uint) error {
	return u.repo.HapusSiswa(id)
}

// --- MATA PELAJARAN ---
func (u *sekolahUsecase) SimpanMapel(mapel *model.MataPelajaran) error {
	return u.repo.TambahMapel(mapel)
}

func (u *sekolahUsecase) DaftarMapel() ([]model.MataPelajaran, error) {
	return u.repo.AmbilSemuaMapel()
}

func (u *sekolahUsecase) DetailMapel(id uint) (model.MataPelajaran, error) {
	return u.repo.AmbilMapelBerdasarkanID(id)
}

func (u *sekolahUsecase) UbahMapel(mapel *model.MataPelajaran) error {
	return u.repo.PerbaruiMapel(mapel)
}

func (u *sekolahUsecase) HapusMapel(id uint) error {
	return u.repo.HapusMapel(id)
}

func (u *sekolahUsecase) BersihkanSemuaData() error {
	return u.repo.HapusSemuaData()
}
