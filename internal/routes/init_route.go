package routes

import (
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handlers.Handlers) {
	authRoute(r, h)
	userRoute(r, h)
	organizationRoute(r, h)
	permissionRoute(r, h)
	roleRoute(r, h)
}
