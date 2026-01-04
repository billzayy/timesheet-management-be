package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.Engine, h *handlers.Handlers) {
	api := r.Group("/api")
	{
		api.POST("/user/create", h.UserHandler.Create)
		api.GET("/user/all", h.UserHandler.GetAll)
		api.GET("/user", h.UserHandler.GetByEmail)
	}
}
