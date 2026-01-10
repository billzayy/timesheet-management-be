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

// Login godoc
//
//	@Summary		Login
//	@Description	Login Timesheet
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginDTO	true	"Send Login Information"
//	@Success		200		{object}	backend.ResponseData
//	@Router			/auth/login [post]
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
