package model

import "time"

type Guru struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Nama         string    `json:"nama"`
	Email        string    `gorm:"unique" json:"email"`
	MataPelajaran string    `json:"mata_pelajaran"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Kelas struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `json:"nama"`
	GuruID    uint      `json:"guru_id"`
	Guru      Guru      `gorm:"foreignKey:GuruID" json:"guru"`
	Siswa     []Siswa   `json:"siswa"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Siswa struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `json:"nama"`
	Email     string    `gorm:"unique" json:"email"`
	KelasID   uint      `json:"kelas_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MataPelajaran struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
