package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/billzayy/timesheet-management-be/config"
	"github.com/billzayy/timesheet-management-be/database"
	"github.com/billzayy/timesheet-management-be/internal/handlers"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
	"github.com/billzayy/timesheet-management-be/internal/routes"
	"github.com/billzayy/timesheet-management-be/internal/services"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)

	if cfg.DBHost == "db" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.Connect(dsn)

	if err != nil {
		log.Fatal("❌ failed to connect database:", err)
	}

	handler := handlers.NewHandlers(
		services.NewServices(
			repositories.NewRepositories(db),
		),
	)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},           // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Exposed headers
		AllowCredentials: true,                                                // Allow credentials (cookies)
		MaxAge:           12 * time.Hour,                                      // Cache duration for preflight requests
	}))

	r.Use(gin.Logger())
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	routes.Register(r, handler)

	r.Run(":" + os.Getenv("REST_PORT"))
}
