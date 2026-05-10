package model

import (
	"time"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Brand       string         `gorm:"size:100;not null" json:"brand"`
	Model       string         `gorm:"size:100;not null" json:"model"`
	PlateNumber string         `gorm:"size:20;unique;not null" json:"plate_number"`
	Status      string         `gorm:"size:20;default:'Available'" json:"status"` // Available/Rented
	Rentals     []Rental       `gorm:"foreignKey:VehicleID" json:"rentals,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Customer struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"size:100;not null" json:"name"`
	IdentityNumber string         `gorm:"size:50;unique;not null" json:"identity_number"`
	PhoneNumber    string         `gorm:"size:20" json:"phone_number"`
	Rentals        []Rental       `gorm:"foreignKey:CustomerID" json:"rentals,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

type Rental struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	VehicleID  uint           `gorm:"not null" json:"vehicle_id"`
	CustomerID uint           `gorm:"not null" json:"customer_id"`
	RentalDate time.Time      `gorm:"not null" json:"rental_date"`
	ReturnDate time.Time      `json:"return_date"`
	TotalCost  float64        `json:"total_cost"`
	Vehicle    Vehicle        `gorm:"foreignKey:VehicleID" json:"vehicle,omitempty"`
	Customer   Customer       `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
