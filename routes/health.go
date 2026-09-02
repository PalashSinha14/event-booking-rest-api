package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/db"
)

// RegisterHealthRoute wires up /healthz: a liveness/readiness check for
// deployment platforms (Docker healthchecks, Render, load balancers) to
// poll. It also pings the database, since a server that's up but can't
// reach Postgres isn't actually healthy.
func RegisterHealthRoute(server *gin.Engine) {
	server.GET("/healthz", func(c *gin.Context) {
		if err := db.DB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
