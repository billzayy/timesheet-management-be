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
