package handler

import (
	"net/http"
	"study/internal/model"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RentalHandler struct {
	usecase usecase.RentalUsecase
}

func NewRentalHandler(u usecase.RentalUsecase) *RentalHandler {
	return &RentalHandler{usecase: u}
}

func (h *RentalHandler) AddVehicle(c *gin.Context) {
	var v model.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Success: false, Message: err.Error()})
		return
	}
	if err := h.usecase.RegisterVehicle(&v); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.Response{Success: true, Message: "Vehicle added", Data: v})
}

func (h *RentalHandler) AddCustomer(c *gin.Context) {
	var cust model.Customer
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Success: false, Message: err.Error()})
		return
	}
	if err := h.usecase.RegisterCustomer(&cust); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.Response{Success: true, Message: "Customer added", Data: cust})
}

func (h *RentalHandler) CreateRental(c *gin.Context) {
	var req struct {
		VehicleID  uint `json:"vehicle_id" binding:"required"`
		CustomerID uint `json:"customer_id" binding:"required"`
		Days       int  `json:"days" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Success: false, Message: err.Error()})
		return
	}
	rental, err := h.usecase.RentVehicle(req.VehicleID, req.CustomerID, req.Days)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.Response{Success: true, Message: "Rental created", Data: rental})
}
