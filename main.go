package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/db"
	"github.com/palashsinha14/go-rest-api/middlewares"
	"github.com/palashsinha14/go-rest-api/routes"
	"github.com/palashsinha14/go-rest-api/models"
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
		email, _ := c.Get("email")
		userId, _ := c.Get("userId")
		c.HTML(200, "dashboard.html", gin.H{
			"email":  email,
			"userId": userId,
		})
	})

	server.GET("/events-page", func(c *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			c.HTML(500, "events.html", gin.H{
				"error": "Could not load events",
			})
			return
		}
		c.HTML(200, "events.html", gin.H{
			"events": events,
		})
	})

	//Posting form data html page
	server.GET("/create-event", middlewares.Authenticate, func(c *gin.Context) {
		c.HTML(200, "create-event.html", nil)
	})

	//Posting form data endpoint
	server.POST("/create-event", middlewares.Authenticate, func(c *gin.Context) {
		name := c.PostForm("name")
		description := c.PostForm("description")
		location := c.PostForm("location")
		datetimeStr := c.PostForm("datetime")
		parsedTime, err := time.Parse("2006-01-02T15:04", datetimeStr)
		if err != nil {
			c.HTML(400, "create-event.html", gin.H{
				"error": "Invalid date format",
			})
			return
		}
		userId := c.GetInt64("userId")
		event := models.Event{
			Name:        name,
			Description: description,
			Location:    location,
			DateTime:    parsedTime,
			UserID:      userId,
		}
		err = event.Save()
		if err != nil {
			c.HTML(500, "create-event.html", gin.H{
				"error": "Could not create event",
			})
			return
		}
		c.Redirect(http.StatusSeeOther, "/events-page")
	})

	//registration for event
	server.GET("/register-page", middlewares.Authenticate, func(c *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			c.String(500, "Error loading events")
			return
		}
		c.HTML(200, "register.html", gin.H{
			"events": events,
		})
	})

	//Logout feature
	server.GET("/logout", func(c *gin.Context) {
		// Clear cookie
		c.SetCookie("token", "", -1, "/", "", true, true)
		// Redirect to login
		c.Redirect(http.StatusSeeOther, "/login-page")
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
