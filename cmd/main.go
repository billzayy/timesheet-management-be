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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/billzayy/timesheet-management-be/docs"
	"github.com/gin-contrib/cors"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)

	db, err := database.Connect(dsn)

	if err != nil {
		log.Fatal("❌ failed to connect database:", err)
	}

	handler := handlers.NewHandlers(
		services.NewServices(
			repositories.NewRepositories(db),
		),
	)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Description = "Timesheet Management API server."
	// docs.SwaggerInfo.Host = os.Getenv("REST_HOST") + ":" + os.Getenv("REST_PORT")
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3001"}, // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},        // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},                  // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                                 // Exposed headers
		AllowCredentials: true,                                                       // Allow credentials (cookies)
		MaxAge:           12 * time.Hour,                                             // Cache duration for preflight requests
	}))

	r.Use(gin.Logger())
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.Register(r, handler)

	r.Run(":" + os.Getenv("REST_PORT"))
}
