package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/palashsinha14/go-rest-api/models"
	"github.com/palashsinha14/go-rest-api/notifier"
)

// registerForEvent godoc
// @Summary      Register for an event
// @Description  Register the authenticated user for an event. Redirects to /my-registrations-page on success - built for the browser, not a JSON API client.
// @Tags         registrations
// @Produce      plain
// @Param        id   path  int  true  "Event ID"
// @Success      303  "Redirects to /my-registrations-page"
// @Failure      400  {string}  string  "invalid event id"
// @Failure      409  {string}  string  "already registered for this event"
// @Failure      500  {string}  string
// @Security     CookieAuth
// @Router       /events/{id}/register [post]
func registerForEvent(c *gin.Context) {
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

	err = event.Register(userId)
	if err != nil {
		// A unique_violation (Postgres error code 23505) means this user
		// already registered for this event - the UNIQUE(event_id, user_id)
		// constraint caught it. Give a clear message instead of a generic
		// 500 for this specific, expected case.
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			c.String(http.StatusConflict, "You are already registered for this event")
			return
		}
		c.String(500, "Could not register")
		return
	}

	// Queue a (simulated) confirmation job on a background worker instead
	// of handling it inline, so this request returns immediately.
	notifier.Enqueue(notifier.ConfirmationJob{
		UserEmail: c.GetString("email"),
		EventName: event.Name,
	})

	// ✅ Redirect instead of JSON
	c.Redirect(http.StatusSeeOther, "/my-registrations-page")
}

/*
func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}
*/
// cancelRegistration godoc
// @Summary      Cancel a registration
// @Description  Cancel the authenticated user's registration for an event
// @Tags         registrations
// @Produce      json
// @Param        id   path      int  true  "Event ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     CookieAuth
// @Router       /events/{id}/register [delete]
func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistrations(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registeration."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registeration Cancelled!"})
}
