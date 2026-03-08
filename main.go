package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mawinkhumdee/flowly-project/backend/config"
	"github.com/mawinkhumdee/flowly-project/backend/database"
	"github.com/mawinkhumdee/flowly-project/backend/handlers"
)

func main() {
	// Load Configuration
	config.LoadConfig()

	// Connect to Database
	database.ConnectDB()

	r := gin.Default()

	// Configure CORS using config.yml
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.AppConfig.Server.FrontendOrigin}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	// Routes
	api := r.Group("/api")
	{
		api.POST("/seed", handlers.SeedStops)
		api.POST("/signup", handlers.Signup)
		api.POST("/login", handlers.Login)

		api.GET("/trips", handlers.GetTrips)
		api.POST("/trips", handlers.CreateTrip)
		api.GET("/trips/:id", handlers.GetTrip)
		api.DELETE("/trips/:id", handlers.DeleteTrip)
		api.PUT("/trips/:id/sharing", handlers.UpdateTripSharing)

		api.GET("/stops", handlers.GetStops)
		api.POST("/stops", handlers.CreateStop)
		api.PUT("/stops/:id", handlers.UpdateStop)
		api.DELETE("/stops/:id", handlers.DeleteStop)
		api.PUT("/stops/reorder", handlers.ReorderStops)
	}

	log.Printf("Server executing on http://localhost%s", config.AppConfig.Server.Port)
	r.Run(config.AppConfig.Server.Port)
}
