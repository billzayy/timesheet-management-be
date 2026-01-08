package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func permissionRoute(r *gin.Engine, h *handlers.Handlers) {
	group := r.Group("/api/permission", middleware.AuthMiddleware())
	{
		group.GET("/all", h.PermissionHandler.GetListPermission)
	}
}
