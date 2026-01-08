package handlers

import (
	"context"
	"net/http"
	"strconv"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/services"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	service services.RoleService
}

func NewRoleHandler(s services.RoleService) *RoleHandler {
	return &RoleHandler{s}
}

func (h *RoleHandler) GetListRole(c *gin.Context) {
	ctx := context.Background()
	var errStr string

	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	data, err := h.service.GetList(ctx, limit, offset)

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

func (h *RoleHandler) GetRoleById(c *gin.Context) {
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

	data, err := h.service.GetRoleById(ctx, int64(id))

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

func (h *RoleHandler) CreateRole(c *gin.Context) {
	ctx := context.Background()
	var errStr string

	var input dto.RoleDTO

	if err := c.BindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	token := backend.GetTokenId(c)

	if err := h.service.CreateRole(ctx, input, token); err != nil {
		errStr = backend.ErrTokenNotFound.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusCreated, backend.ResponseData{
		Result:  nil,
		Success: true,
		Error:   nil,
	})
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	ctx := context.Background()
	var errStr string

	input := c.Query("id")

	id, err := strconv.Atoi(input)

	if err != nil {
		errStr := err.Error()

		c.JSON(http.StatusBadGateway, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	if err := h.service.DeleteRole(ctx, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result:  nil,
		Success: true,
		Error:   nil,
	})
}
