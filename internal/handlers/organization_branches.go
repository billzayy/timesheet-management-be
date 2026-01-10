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

// Gett All Branches godoc
//
//	@Summary		Get All Branches
//	@Description	Get All Branches Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			limit	query	string	true	"Limit pagination"
//	@Param			offset	query	string	true	"Offset pagination"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/branch/all [get]
func (s *OrganizeHandler) GetAllBranches(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	data, err := s.service.GetAllBranches(ctx, limit, offset)

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

// Create Branch godoc
//
//	@Summary		Create Branch
//	@Description	Create Branch
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.BranchDTO	true	"Create Branch Information"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/branch/create [post]
func (s *OrganizeHandler) CreateBranch(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	var input dto.BranchDTO

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

	if err := s.service.CreateBranch(ctx, input, id); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Create Branch successful",
		Success: true,
		Error:   nil,
	})
}

// Update Branch godoc
//
//	@Summary		Update Branch
//	@Description	Update Branch Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.BranchDTO	true	"Update Branch Information"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/branch/update [put]
func (s *OrganizeHandler) UpdateBranch(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	var input dto.BranchDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if err := s.service.UpdateBranch(ctx, input); err != nil {
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
		Result:  "Update Branch successful",
		Success: true,
		Error:   nil,
	})
}

// Delete Branch godoc
//
//	@Summary		Delete Branch
//	@Description	Delete Branch
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id	query	string	true	"Branch Id"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/branch/delete [delete]
func (s *OrganizeHandler) DeleteBranch(c *gin.Context) {
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

	if err := s.service.DeleteBranch(ctx, int64(id)); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Delete Branch successful",
		Success: true,
		Error:   nil,
	})

}
