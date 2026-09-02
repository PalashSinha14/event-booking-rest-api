package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/middlewares"
	"github.com/palashsinha14/go-rest-api/models"
)

// RegisterMyEventsRoutes wires up the "My Events" and "My Registrations"
// HTML pages: events the logged-in user created, and events they've
// registered for, each with an action (delete / cancel) on the page.
func RegisterMyEventsRoutes(server *gin.Engine) {

	// My Events page - events created by the logged-in user
	server.GET("/my-events-page", middlewares.Authenticate, func(c *gin.Context) {
		userId := c.GetInt64("userId")
		events, err := models.GetEventsByUserID(userId)
		if err != nil {
			c.String(500, "Error loading your events")
			return
		}
		c.HTML(200, "my-events.html", gin.H{
			"events": events,
		})
	})

	// Delete an event from the My Events page
	server.POST("/my-events/:id/delete", middlewares.Authenticate, func(c *gin.Context) {
		userId := c.GetInt64("userId")
		eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.String(400, "Invalid event ID")
			return
		}
		event, err := models.GetEventByID(eventId)
		if err != nil {
			c.String(500, "Event not found")
			return
		}
		if event.UserID != userId {
			c.String(http.StatusUnauthorized, "Not authorized to delete this event")
			return
		}
		if err := event.Delete(); err != nil {
			c.String(500, "Could not delete event")
			return
		}
		c.Redirect(http.StatusSeeOther, "/my-events-page")
	})

	// My Registrations page - events the logged-in user has registered for
	server.GET("/my-registrations-page", middlewares.Authenticate, func(c *gin.Context) {
		userId := c.GetInt64("userId")
		events, err := models.GetRegisteredEventsByUserID(userId)
		if err != nil {
			c.String(500, "Error loading your registrations")
			return
		}
		c.HTML(200, "my-registrations.html", gin.H{
			"events": events,
		})
	})

	// Cancel a registration from the My Registrations page
	server.POST("/my-registrations/:id/cancel", middlewares.Authenticate, func(c *gin.Context) {
		userId := c.GetInt64("userId")
		eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.String(400, "Invalid event ID")
			return
		}
		var event models.Event
		event.ID = eventId
		if err := event.CancelRegistrations(userId); err != nil {
			c.String(500, "Could not cancel registration")
			return
		}
		c.Redirect(http.StatusSeeOther, "/my-registrations-page")
	})
}
