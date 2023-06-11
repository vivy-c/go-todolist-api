package handler

import (
	"go-simple-template/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ToDoUpdate(c echo.Context) error {
	// Get the book ID from the URL parameter
	id := c.Param("id")

	// Bind the request body to the update request DTO
	dataReq := dto.ToDoUpdateReq{}
	err := c.Bind(&dataReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, dto.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "cannot read payload",
		})
	}

	// Call the service method to update the todo
	err = h.service.ToDoUpdate(id, dataReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Message: "todo updated successfully",
	})
}
