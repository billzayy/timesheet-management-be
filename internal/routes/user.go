package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.Engine, h *handlers.Handlers) {
	api := r.Group("/api/user", middleware.AuthMiddleware())
	{
		api.POST("/create", h.UserHandler.Create)
		api.GET("/all", h.UserHandler.GetAll)
		api.GET("/:id", h.UserHandler.GetById)
		api.GET("/by-email", h.UserHandler.GetByEmail)
		api.DELETE("/", h.UserHandler.Delete)
	}
}
