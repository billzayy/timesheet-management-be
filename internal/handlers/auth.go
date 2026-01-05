package handlers

import (
	"context"
	"net/http"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(s services.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Login(c *gin.Context) {
	ctx := context.Background()
	c.Header("Content-Type", "application/json")

	var input dto.LoginDTO
	var errStr string

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	data, err := h.service.Login(ctx, &input)

	if err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result:  data,
		Success: true,
		Error:   nil,
	})
}
