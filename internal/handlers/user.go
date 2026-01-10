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
	"github.com/google/uuid"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{s}
}

// Create User godoc
//
//	@Summary		Create User
//	@Description	Create User Information
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RequestUserDTO	true	"Send Request Create User"
//	@Success		200		{object}	backend.ResponseData
//	@Failure		400		{object}	backend.ResponseData
//	@Router			/user/create [post]
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

// Get All Users godoc
//
//	@Summary		Get All
//	@Description	Gett All Users Information
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			limit	query		string	true	"Limit Get"
//	@Param			offset	query		string	true	"Offset for pagination"
//	@Success		200		{object}	backend.ResponseData
//	@Router			/user/all [get]
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

// Get User By Email godoc
//
//	@Summary		Get User By Email
//	@Description	Gett User Information By Email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			email	query		string	true	"Email User"
//	@Success		200		{object}	backend.ResponseData
//	@Router			/user/by-email [get]
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

// Get User By Id godoc
//
//	@Summary		Get User By Id
//	@Description	Get User Information By Id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Id User"
//	@Success		200	{object}	backend.ResponseData
//	@Router			/user/{id} [get]
func (h *UserHandler) GetById(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	id := c.Param("id")

	if id == "" {
		errStr = "input must not be empty"

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  "Error",
			Success: false,
			Error:   &errStr,
		})
		return
	}

	data, err := h.service.GetById(ctx, uuid.MustParse(id))

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

// Delete User By Id godoc
//
//	@Summary		Delete User By Id
//	@Description	Delete User Information By Id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	query		string	true	"Id User"
//	@Success		200	{object}	backend.ResponseData
//	@Router			/user/delete [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	id := c.Query("id")

	if id == "" {
		errStr = "input must not be empty"

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  "Error",
			Success: false,
			Error:   &errStr,
		})
		return
	}

	err := h.service.Delete(ctx, id)

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
