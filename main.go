package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/db"
	"github.com/palashsinha14/go-rest-api/middlewares"
	"github.com/palashsinha14/go-rest-api/notifier"
	"github.com/palashsinha14/go-rest-api/routes"
)

func main() {

	db.InitDB()

	notifier.StartWorkers(3)
	middlewares.StartRateLimiterCleanup()

	server := gin.Default()

	server.Static("/static", "./frontend")
	server.LoadHTMLGlob("frontend/*.html")

	routes.RegisterPageRoutes(server)
	routes.RegisterRoutes(server)
	routes.RegisterMyEventsRoutes(server)

	// Render / Docker dynamic port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)

	server.Run(":" + port)
}
