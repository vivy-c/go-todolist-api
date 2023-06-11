package handler

import (
	"go-simple-template/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ToDoAll(c echo.Context) error {
	ToDo, err := h.service.ToDoAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Message: "To do fetched successfully",
		Data:    ToDo, // Assuming ToDo is a slice of book objects
	})
}
