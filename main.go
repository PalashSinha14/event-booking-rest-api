package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/db"
	"github.com/palashsinha14/go-rest-api/middlewares"
	"github.com/palashsinha14/go-rest-api/notifier"
	"github.com/palashsinha14/go-rest-api/routes"
)

// @title           Event Booking REST API
// @version         1.0
// @description     A production-ready RESTful backend service for creating, browsing, and registering for events. Built with Go, Gin, and PostgreSQL.
// @contact.name    Palash Sinha
// @host            event-booking-rest-api.onrender.com
// @BasePath        /
// @schemes         https

// @securityDefinitions.apikey  CookieAuth
// @in                          cookie
// @name                        token
// @description                 JWT issued by POST /login, stored as an httpOnly cookie named "token".
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
	routes.RegisterHealthRoute(server)
	routes.RegisterSwaggerRoute(server)

	// Render / Docker dynamic port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
	}

	// Run the server in the background so the main goroutine is free to
	// wait for a shutdown signal below.
	go func() {
		log.Println("Server running on port:", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Block until an interrupt/terminate signal arrives (e.g. Ctrl+C
	// locally, or SIGTERM from Docker/Render when stopping the container).
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	<-ctx.Done()
	stop()

	log.Println("Shutdown signal received, finishing in-flight requests...")

	// Give in-flight requests up to 10 seconds to finish before forcing
	// the connections closed.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shut down:", err)
	}

	db.DB.Close()

	log.Println("Server exited cleanly")
}
