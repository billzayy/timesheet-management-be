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

// Get All Levels godoc
//
//	@Summary		Get All Levels
//	@Description	Get All Levels Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			limit	query	string	true	"Limit pagination"
//	@Param			offset	query	string	true	"Offset pagination"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/level/all [get]
func (s *OrganizeHandler) GetAllLevels(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	data, err := s.service.GetAllLevels(ctx, limit, offset)

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

// Create Level godoc
//
//	@Summary		Create Level
//	@Description	Create Level
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.LevelDTO	true	"Create Level Information"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/level/create [post]
func (s *OrganizeHandler) CreateLevel(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	var input dto.LevelDTO

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

	if err := s.service.CreateLevel(ctx, input, id); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Create level successful",
		Success: true,
		Error:   nil,
	})
}

// Update Level godoc
//
//	@Summary		Update Level
//	@Description	Update Level
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.LevelDTO	true	"Create Level Information"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/level/update [put]
func (s *OrganizeHandler) UpdateLevel(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	var input dto.LevelDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if err := s.service.UpdateLevel(ctx, input); err != nil {
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
		Result:  "Update level successful",
		Success: true,
		Error:   nil,
	})
}

// Delete Level godoc
//
//	@Summary		Delete Level
//	@Description	Delete Level
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id	query	string	true	"Delete Level Information"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/level/delete [delete]
func (s *OrganizeHandler) DeleteLevel(c *gin.Context) {
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

	if err := s.service.DeleteLevel(ctx, int64(id)); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Delete level successful",
		Success: true,
		Error:   nil,
	})

}
