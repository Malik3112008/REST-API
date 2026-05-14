package repository

import (
	"errors"
	"gorm.io/gorm"
	"study/internal/model"
)

type SekolahRepository interface {
	
	TambahGuru(guru *model.Guru) error
	AmbilSemuaGuru() ([]model.Guru, error)
	AmbilGuruBerdasarkanID(id uint) (model.Guru, error)
	PerbaruiGuru(guru *model.Guru) error
	HapusGuru(id uint) error

	
	TambahKelas(kelas *model.Kelas) error
	AmbilSemuaKelas() ([]model.Kelas, error)
	AmbilKelasBerdasarkanID(id uint) (model.Kelas, error)
	PerbaruiKelas(kelas *model.Kelas) error
	HapusKelas(id uint) error

	
	TambahSiswa(siswa *model.Siswa) error
	AmbilSemuaSiswa() ([]model.Siswa, error)
	AmbilSiswaBerdasarkanID(id uint) (model.Siswa, error)
	PerbaruiSiswa(siswa *model.Siswa) error
	HapusSiswa(id uint) error

	
	TambahMapel(mapel *model.MataPelajaran) error
	AmbilSemuaMapel() ([]model.MataPelajaran, error)
	AmbilMapelBerdasarkanID(id uint) (model.MataPelajaran, error)
	PerbaruiMapel(mapel *model.MataPelajaran) error
	HapusMapel(id uint) error
}

type sekolahRepository struct {
	db *gorm.DB
}

func NewSekolahRepository(db *gorm.DB) SekolahRepository {
	return &sekolahRepository{db}
}


func (r *sekolahRepository) TambahGuru(guru *model.Guru) error {
	if err := r.db.Create(guru).Error; err != nil {
		return err
	}
	return r.db.Preload("MataPelajaran").First(guru, guru.ID).Error
}

func (r *sekolahRepository) AmbilSemuaGuru() ([]model.Guru, error) {
	var daftarGuru []model.Guru
	err := r.db.Preload("MataPelajaran").Find(&daftarGuru).Error
	return daftarGuru, err
}

func (r *sekolahRepository) AmbilGuruBerdasarkanID(id uint) (model.Guru, error) {
	var guru model.Guru
	err := r.db.Preload("MataPelajaran").First(&guru, id).Error
	return guru, err
}

func (r *sekolahRepository) PerbaruiGuru(guru *model.Guru) error {
	if err := r.db.Save(guru).Error; err != nil {
		return err
	}
	return r.db.Preload("MataPelajaran").First(guru, guru.ID).Error
}

func (r *sekolahRepository) HapusGuru(id uint) error {
	result := r.db.Delete(&model.Guru{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("guru tidak ditemukan")
	}
	return nil
}


func (r *sekolahRepository) TambahKelas(kelas *model.Kelas) error {
	if err := r.db.Create(kelas).Error; err != nil {
		return err
	}
	return r.db.Preload("Guru.MataPelajaran").Preload("Siswa").First(kelas, kelas.ID).Error
}

func (r *sekolahRepository) AmbilSemuaKelas() ([]model.Kelas, error) {
	var daftarKelas []model.Kelas
	err := r.db.Preload("Guru.MataPelajaran").Preload("Siswa").Find(&daftarKelas).Error
	return daftarKelas, err
}

func (r *sekolahRepository) AmbilKelasBerdasarkanID(id uint) (model.Kelas, error) {
	var kelas model.Kelas
	err := r.db.Preload("Guru.MataPelajaran").Preload("Siswa").First(&kelas, id).Error
	return kelas, err
}

func (r *sekolahRepository) PerbaruiKelas(kelas *model.Kelas) error {
	if err := r.db.Save(kelas).Error; err != nil {
		return err
	}
	return r.db.Preload("Guru.MataPelajaran").Preload("Siswa").First(kelas, kelas.ID).Error
}

func (r *sekolahRepository) HapusKelas(id uint) error {
	result := r.db.Delete(&model.Kelas{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("kelas tidak ditemukan")
	}
	return nil
}


func (r *sekolahRepository) TambahSiswa(siswa *model.Siswa) error {
	if err := r.db.Create(siswa).Error; err != nil {
		return err
	}
	return r.db.Preload("Kelas.Guru.MataPelajaran").First(siswa, siswa.ID).Error
}

func (r *sekolahRepository) AmbilSemuaSiswa() ([]model.Siswa, error) {
	var daftarSiswa []model.Siswa
	err := r.db.Preload("Kelas.Guru.MataPelajaran").Find(&daftarSiswa).Error
	return daftarSiswa, err
}

func (r *sekolahRepository) AmbilSiswaBerdasarkanID(id uint) (model.Siswa, error) {
	var siswa model.Siswa
	err := r.db.Preload("Kelas.Guru.MataPelajaran").First(&siswa, id).Error
	return siswa, err
}

func (r *sekolahRepository) PerbaruiSiswa(siswa *model.Siswa) error {
	if err := r.db.Save(siswa).Error; err != nil {
		return err
	}
	return r.db.Preload("Kelas.Guru.MataPelajaran").First(siswa, siswa.ID).Error
}

func (r *sekolahRepository) HapusSiswa(id uint) error {
	result := r.db.Delete(&model.Siswa{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("siswa tidak ditemukan")
	}
	return nil
}


func (r *sekolahRepository) TambahMapel(mapel *model.MataPelajaran) error {
	return r.db.Create(mapel).Error
}

func (r *sekolahRepository) AmbilSemuaMapel() ([]model.MataPelajaran, error) {
	var daftarMapel []model.MataPelajaran
	err := r.db.Find(&daftarMapel).Error
	return daftarMapel, err
}

func (r *sekolahRepository) AmbilMapelBerdasarkanID(id uint) (model.MataPelajaran, error) {
	var mapel model.MataPelajaran
	err := r.db.First(&mapel, id).Error
	return mapel, err
}

func (r *sekolahRepository) PerbaruiMapel(mapel *model.MataPelajaran) error {
	return r.db.Save(mapel).Error
}

func (r *sekolahRepository) HapusMapel(id uint) error {
	result := r.db.Delete(&model.MataPelajaran{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("mata pelajaran tidak ditemukan")
	}
	return nil
}
