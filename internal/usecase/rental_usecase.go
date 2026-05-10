package usecase

import (
	"errors"
	"study/internal/model"
	"study/internal/repository"
	"time"
)

type RentalUsecase interface {
	RegisterVehicle(vehicle *model.Vehicle) error
	RegisterCustomer(customer *model.Customer) error
	RentVehicle(vehicleID, customerID uint, days int) (*model.Rental, error)
}

type rentalUsecase struct {
	repo repository.RentalRepository
}

func NewRentalUsecase(repo repository.RentalRepository) RentalUsecase {
	return &rentalUsecase{repo: repo}
}

func (u *rentalUsecase) RegisterVehicle(v *model.Vehicle) error {
	return u.repo.CreateVehicle(v)
}

func (u *rentalUsecase) RegisterCustomer(c *model.Customer) error {
	return u.repo.CreateCustomer(c)
}

func (u *rentalUsecase) RentVehicle(vID, cID uint, days int) (*model.Rental, error) {
	vehicle, err := u.repo.GetVehicleByID(vID)
	if err != nil {
		return nil, err
	}

	if vehicle.Status != "Available" {
		return nil, errors.New("vehicle is not available for rent")
	}

	rental := &model.Rental{
		VehicleID:  vID,
		CustomerID: cID,
		RentalDate: time.Now(),
		ReturnDate: time.Now().AddDate(0, 0, days),
		TotalCost:  float64(days) * 100000, // Flat rate example
	}

	if err := u.repo.CreateRental(rental); err != nil {
		return nil, err
	}

	return rental, nil
}
