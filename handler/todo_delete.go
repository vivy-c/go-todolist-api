package handler

import (
	"go-simple-template/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ToDoDelete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
	}

	err = h.service.ToDoDelete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Message: "Book deleted successfully",
	})
}
