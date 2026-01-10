package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Get All Positions godoc
//
//	@Summary		Get All Positions
//	@Description	Get All Positions Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			limit	query	string	true	"Limit pagination"
//	@Param			offset	query	string	true	"Offset pagination"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/position/all [get]
func (s *OrganizeHandler) GetAllPositions(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	data, err := s.service.GetAllPositions(ctx, limit, offset)

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

// Create Position godoc
//
//	@Summary		Create Position
//	@Description	Create Position Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.PositionDTO	true	"Request Body"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/position/create [post]
func (s *OrganizeHandler) CreatePosition(c *gin.Context) {
	ctx := context.Background()

	var errStr string

	var input dto.PositionDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	value, ok := c.Get("token")

	if !ok {
		errStr = backend.ErrTokenNotFound.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	var id uuid.UUID

	switch t := value.(type) {
	case string:
		id = uuid.MustParse(t)
	default:
		id = uuid.Nil
	}

	if err := s.service.CreatePosition(ctx, input, id); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Create position successful",
		Success: true,
		Error:   nil,
	})

}

// Update Position godoc
//
//	@Summary		Update Position
//	@Description	Update Position Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.PositionDTO	true	"Request Body"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/position/update [put]
func (s *OrganizeHandler) UpdatePosition(c *gin.Context) {
	ctx := context.Background()

	var errStr string
	var input dto.PositionDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if err := s.service.UpdatePosition(ctx, input); err != nil {
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
		Result:  "Update position successful",
		Success: true,
		Error:   nil,
	})
}

// Delete Position godoc
//
//	@Summary		Delete Position
//	@Description	Delete Position Information
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Param			id	query	string	true	"Position Id"
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/position/delete [delete]
func (s *OrganizeHandler) DeletePosition(c *gin.Context) {
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

	if err := s.service.DeletePosition(ctx, int64(id)); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  "Delete position successful",
		Success: true,
		Error:   nil,
	})

}
