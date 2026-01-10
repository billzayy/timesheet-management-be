package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/gin-gonic/gin"
)

// Get All User Types godoc
//
//	@Summary		Get All User Types
//	@Description	Get All User Types Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			limit	query	string	true	"Limit pagination"
//	@Param			offset	query	string	true	"Offset pagination"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/user-type/all [get]
func (s *OrganizeHandler) GetAllUserTypes(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	data, err := s.service.GetAllUserTypes(ctx, limit, offset)

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
		Result: backend.RespList{
			TotalCount: len(data),
			Items:      data,
		},
		Success: true,
		Error:   nil,
	})
}

// Create User Types godoc
//
//	@Summary		Create User Types
//	@Description	Create User Types Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.UserTypeDTO	true	"UserType Information"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/user-type/create [post]
func (s *OrganizeHandler) CreateUserType(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	var input dto.UserTypeDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	id := backend.GetTokenId(c)

	if err := s.service.CreateUserType(ctx, input, id); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Create User Type successful",
		Success: true,
		Error:   nil,
	})
}

// Update User Types godoc
//
//	@Summary		Update User Types
//	@Description	Update User Types Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.UserTypeDTO	true	"UserType Information"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/user-type/update [put]
func (s *OrganizeHandler) UpdateUserType(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	var input dto.UserTypeDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if err := s.service.UpdateUserType(ctx, input); err != nil {
		errStr = err.Error()

		if errors.Is(err, backend.ErrPositionNotFound) {
			c.JSON(http.StatusNotFound, backend.ResponseData{
				Result:  nil,
				Success: false,
				Error:   &errStr,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result:  "Update User Type successful",
		Success: true,
		Error:   nil,
	})
}

// Delete User Types godoc
//
//	@Summary		Delete User Types
//	@Description	Delete User Types Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id	query	string	true	"UserType Id"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/user-type/delete [delete]
func (s *OrganizeHandler) DeleteUserType(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if id == 0 {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if err := s.service.DeleteUserType(ctx, int64(id)); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Delete User type successful",
		Success: true,
		Error:   nil,
	})

}
