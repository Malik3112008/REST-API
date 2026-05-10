package repository

import (
	"study/internal/model"
	"gorm.io/gorm"
)

type RentalRepository interface {
	CreateVehicle(vehicle *model.Vehicle) error
	GetVehicleByID(id uint) (*model.Vehicle, error)
	UpdateVehicle(vehicle *model.Vehicle) error
	CreateCustomer(customer *model.Customer) error
	CreateRental(rental *model.Rental) error
}

type rentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &rentalRepository{db: db}
}

func (r *rentalRepository) CreateVehicle(v *model.Vehicle) error {
	return r.db.Create(v).Error
}

func (r *rentalRepository) GetVehicleByID(id uint) (*model.Vehicle, error) {
	var v model.Vehicle
	err := r.db.First(&v, id).Error
	return &v, err
}

func (r *rentalRepository) UpdateVehicle(v *model.Vehicle) error {
	return r.db.Save(v).Error
}

func (r *rentalRepository) CreateCustomer(c *model.Customer) error {
	return r.db.Create(c).Error
}

func (r *rentalRepository) CreateRental(re *model.Rental) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(re).Error; err != nil {
			return err
		}
		return tx.Model(&model.Vehicle{}).Where("id = ?", re.VehicleID).Update("status", "Rented").Error
	})
}
