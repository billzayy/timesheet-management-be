package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func authRoute(r *gin.Engine, h *handlers.Handlers) {
	api := r.Group("/api/auth")
	{
		api.POST("/login", h.AuthHandler.Login)
		api.POST("/refresh-token", h.AuthHandler.RefreshToken)
	}

	r.GET("/api/user-config", middleware.AuthMiddleware(), h.AuthHandler.UserConfig)
}
