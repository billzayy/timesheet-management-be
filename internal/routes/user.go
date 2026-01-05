package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.Engine, h *handlers.Handlers) {
	api := r.Group("/api", middleware.AuthMiddleware())
	{
		api.POST("/user/create", h.UserHandler.Create)
		api.GET("/user/all", h.UserHandler.GetAll)
		api.GET("/user", h.UserHandler.GetByEmail)
		api.DELETE("/user", h.UserHandler.Delete)
	}
}
