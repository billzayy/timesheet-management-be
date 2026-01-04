package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handlers.Handlers) {
	userRoute(r, h)
}
