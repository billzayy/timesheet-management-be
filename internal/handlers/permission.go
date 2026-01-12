package handlers

import (
	"net/http"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/services"
	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	service services.PermissionService
}

func NewPermissionHandler(s services.PermissionService) *PermissionHandler {
	return &PermissionHandler{s}
}

// Get List Permission godoc
//
//	@Summary		Get List Permission
//	@Description	Get List Permission
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	backend.ResponseData
//	@Router			/permission/all [get]
func (h *PermissionHandler) GetListPermission(c *gin.Context) {
	var errStr string

	data, err := h.service.GeAllPermission()

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
