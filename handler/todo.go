package handler

import (
	"go-simple-template/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ToDoCreate(c echo.Context) error {
	dataReq := dto.ToDoCreateReq{}
	err := c.Bind(&dataReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, dto.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "cannot read payload",
		})
	}

	err = h.service.ToDoCreate(dataReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.ApiResponse{
		Code:    http.StatusCreated,
		Message: "todo created successfully",
	})
}
