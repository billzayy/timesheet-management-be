package handlers

import (
	"context"
	"errors"
	"net/http"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) Create(c *gin.Context) {
	ctx := context.Background()

	var inputUser dto.RequestUserDTO
	var errStr string

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		errStr = err.Error()
		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if err := h.service.CreateUser(ctx, &inputUser); err != nil {
		errStr = err.Error()
		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Created User Successful",
		Success: true,
		Error:   nil,
	})
}

func (h *UserHandler) GetAll(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	limit := c.Query("limit")
	offset := c.Query("offset")

	users, err := h.service.GetAllUsers(ctx, limit, offset)

	if err != nil {
		errStr = err.Error()
		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  "Error",
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result: backend.RespList{
			TotalCount: len(users),
			Items:      users,
		},
		Success: true,
		Error:   nil,
	})
}

func (h *UserHandler) GetByEmail(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	email := c.Query("email")

	if email == "" {
		errStr = "input must not be empty"

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  "Error",
			Success: false,
			Error:   &errStr,
		})
		return
	}

	// Validate email format
	if validator.New().Var(email, "email") != nil {
		errStr = "failed to validate email format"

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  "Error",
			Success: false,
			Error:   &errStr,
		})
		return
	}

	data, err := h.service.GetByEmail(ctx, email)

	if err != nil {
		errStr := err.Error()

		switch {
		case errors.Is(err, backend.ErrUserNotFound):
			c.JSON(http.StatusNotFound, backend.ResponseData{
				Result:  "Error",
				Success: false,
				Error:   &errStr,
			})
		default:
			c.JSON(http.StatusInternalServerError, backend.ResponseData{
				Result:  "Error",
				Success: false,
				Error:   &errStr,
			})
		}
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result:  data,
		Success: true,
		Error:   nil,
	})
}

func (h *UserHandler) Delete(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	email := c.Query("email")

	if email == "" {
		errStr = "input must not be empty"

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  "Error",
			Success: false,
			Error:   &errStr,
		})
		return
	}

	// Validate email format
	if validator.New().Var(email, "email") != nil {
		errStr = "failed to validate email format"

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  "Error",
			Success: false,
			Error:   &errStr,
		})
		return
	}

	err := h.service.DeleteByEmail(ctx, email)

	if err != nil {
		errStr := err.Error()

		switch {
		case errors.Is(err, backend.ErrUserNotFound):
			c.JSON(http.StatusNotFound, backend.ResponseData{
				Result:  "Error",
				Success: false,
				Error:   &errStr,
			})
		default:
			c.JSON(http.StatusInternalServerError, backend.ResponseData{
				Result:  "Error",
				Success: false,
				Error:   &errStr,
			})
		}
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result:  "Delete user successful",
		Success: true,
		Error:   nil,
	})
}
