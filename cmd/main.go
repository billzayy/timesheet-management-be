package main

import (
	"fmt"
	"log"

	"github.com/billzayy/timesheet-management-be/config"
	"github.com/billzayy/timesheet-management-be/database"
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
	"github.com/billzayy/timesheet-management-be/internal/routes"
	"github.com/billzayy/timesheet-management-be/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=Asia/Shanghai",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)

	db := database.Connect(dsn)

	handler := handlers.NewHandlers(
		services.NewServices(
			repositories.NewRepositories(db),
		),
	)

	r := gin.Default()
	routes.Register(r, handler)

	r.Run(":8080")
}
