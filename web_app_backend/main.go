package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julien-mrty/Web_app_jump_higher/web_app_backend/api"
	"github.com/julien-mrty/Web_app_jump_higher/web_app_backend/database"
	_ "github.com/julien-mrty/Web_app_jump_higher/web_app_backend/docs" // Swagger documentation initialization
	"github.com/julien-mrty/Web_app_jump_higher/web_app_backend/middleware"
	"github.com/julien-mrty/Web_app_jump_higher/web_app_backend/services"
)

// @title Web App Jump Higher API
// @version 0.1
// @description This is the API documentation for the Web App Jump Higher backend.
// @termsOfService http://Nanmarkalkab.com/terms/

// @contact.name API La Mi
// @contact.url http://xoxo.com/support
// @contact.email support@xoxo.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /

func main() {
	// Initialize logging and validation services
	services.InitLogger()
	services.InitValidator()

	// Create a Gin router instance
	r := gin.Default()

	// Apply global middleware
	r.Use(middleware.CORSMiddleware())

	// Initialize the database
	for i := 0; i < 10; i++ {
		err := database.ConnectDB()
		if err == nil {
			break
		}
		log.Printf("Can not reach database. Retrying in 5 seconds... (%d/10)", i+1)
		time.Sleep(5 * time.Second)
	}

	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Setup application routes
	api.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
