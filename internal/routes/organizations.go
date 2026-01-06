package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func organizationRoute(r *gin.Engine, h *handlers.Handlers) {
	positionGroup := r.Group("/api/position", middleware.AuthMiddleware())
	{
		positionGroup.GET("/all", h.OrganizeHandler.GetAllPositions)
		positionGroup.POST("/create", h.OrganizeHandler.CreatePosition)
		positionGroup.PUT("/update", h.OrganizeHandler.UpdatePosition)
		positionGroup.DELETE("/delete", h.OrganizeHandler.DeletePosition)
	}

	levelGroup := r.Group("/api/level", middleware.AuthMiddleware())
	{
		levelGroup.GET("/all", h.OrganizeHandler.GetAllLevels)
		levelGroup.POST("/create", h.OrganizeHandler.CreateLevel)
		levelGroup.PUT("/update", h.OrganizeHandler.UpdateLevel)
		levelGroup.DELETE("/delete", h.OrganizeHandler.DeleteLevel)
	}

	userTypeGroup := r.Group("/api/user-type", middleware.AuthMiddleware())
	{
		userTypeGroup.GET("/all", h.OrganizeHandler.GetAllUserTypes)
		userTypeGroup.POST("/create", h.OrganizeHandler.CreateUserType)
		userTypeGroup.PUT("/update", h.OrganizeHandler.UpdateUserType)
		userTypeGroup.DELETE("/delete", h.OrganizeHandler.DeleteUserType)
	}
}
