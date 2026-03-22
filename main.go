package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/db"
	"github.com/palashsinha14/go-rest-api/routes"
	"github.com/palashsinha14/go-rest-api/middlewares"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.Static("/static", "./frontend")
	server.LoadHTMLGlob("frontend/*.html")

	// Frontend routes
	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	server.GET("/login-page", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	server.GET("/signup-page", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})

/*	server.GET("/dashboard", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", nil)
	})*/
/*
	server.GET("/dashboard", middlewares.AuthMiddlewareHTML, func(c *gin.Context) {
	c.HTML(200, "dashboard.html", nil)
	})
*/
	server.GET("/dashboard", middlewares.Authenticate, func(c *gin.Context) {
		c.HTML(200, "dashboard.html", nil)
	})
	routes.RegisterRoutes(server)

	// Render / Docker dynamic port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)

	server.Run(":" + port)
}

/*
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/db"
	"github.com/palashsinha14/go-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.Static("/static", "./frontend")
	server.LoadHTMLGlob("frontend/*.html")

	// Frontend routes
	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	server.GET("/login-page", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	server.GET("/signup-page", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})

	routes.RegisterRoutes(server)

	// Render requires dynamic port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start server
	server.Run(":" + port)
}*/

/*
package main

//import "fmt"

import (
	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/db"
	"github.com/palashsinha14/go-rest-api/routes"
)

func main() {
	//fmt.Println("Hello World")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
*/
