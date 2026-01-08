package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func roleRoute(r *gin.Engine, h *handlers.Handlers) {
	roleGroup := r.Group("/api/role", middleware.AuthMiddleware())
	{
		roleGroup.GET("/all", h.RoleHandler.GetListRole)
		roleGroup.POST("/create", h.RoleHandler.CreateRole)
		roleGroup.DELETE("/delete", h.RoleHandler.DeleteRole)
	}
}
