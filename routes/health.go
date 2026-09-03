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
	server.GET("/healthz", healthz)
}

// healthz godoc
// @Summary      Health check
// @Description  Liveness/readiness check - also pings the database
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      503  {object}  map[string]string
// @Router       /healthz [get]
func healthz(c *gin.Context) {
	if err := db.DB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
