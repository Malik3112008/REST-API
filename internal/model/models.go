package model

import "time"

type Guru struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	Nama            string        `json:"nama"`
	Email           string        `gorm:"unique" json:"email"`
	MataPelajaranID uint          `json:"mata_pelajaran_id"`
	MataPelajaran   MataPelajaran `gorm:"foreignKey:MataPelajaranID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"mata_pelajaran"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type Kelas struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `json:"nama"`
	GuruID    uint      `json:"guru_id"`
	Guru      Guru      `gorm:"foreignKey:GuruID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"guru"`
	Siswa     []Siswa   `gorm:"foreignKey:KelasID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"siswa"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Siswa struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `json:"nama"`
	Email     string    `gorm:"unique" json:"email"`
	KelasID   uint      `json:"kelas_id"`
	Kelas     Kelas     `gorm:"foreignKey:KelasID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"kelas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MataPelajaran struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"unique" json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
