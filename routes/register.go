package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/models"
	"github.com/palashsinha14/go-rest-api/notifier"
)

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
