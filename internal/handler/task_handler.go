package handler

import (
	"net/http"
	"study/internal/model"
	"study/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase usecase.TaskUsecase
}

func NewTaskHandler(u usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: u}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Input tidak valid",
			Errors:  err.Error(),
		})
		return
	}

	if err := h.usecase.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Gagal membuat tugas",
			Errors:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Success: true,
		Message: "Tugas berhasil dibuat",
		Data:    task,
	})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "ID tidak sesuai",
		})
		return
	}

	task, err := h.usecase.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Success: false,
			Message: "Tugas tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Data:    task,
	})
}

func (h *TaskHandler) UpdateTaskStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "ID tidak sesuai",
		})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Status wajib diisi",
		})
		return
	}

	if err := h.usecase.UpdateTaskStatus(uint(id), input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Gagal memperbarui status",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Status berhasil diperbarui",
	})
}
