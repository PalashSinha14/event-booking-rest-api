package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/models"
)

// getEvents godoc
// @Summary      List all events
// @Description  Fetch every event in the system
// @Tags         events
// @Produce      json
// @Success      200  {array}   models.Event
// @Failure      500  {object}  map[string]string
// @Router       /events [get]
func getEvents(c *gin.Context) {
	events,err:=models.GetAllEvents()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	c.JSON(http.StatusOK, events)
}

// getEvent godoc
// @Summary      Get a single event
// @Description  Fetch one event by its ID
// @Tags         events
// @Produce      json
// @Param        id   path      int  true  "Event ID"
// @Success      200  {object}  models.Event
// @Failure      400  {object}  map[string]string  "invalid event id"
// @Failure      500  {object}  map[string]string
// @Router       /events/{id} [get]
func getEvent(c *gin.Context){
	eventId, err:=strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err:=models.GetEventByID(eventId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event."})
		return
	}
	c.JSON(http.StatusOK,event)
}

// createEvent godoc
// @Summary      Create an event
// @Description  Create a new event. The authenticated user becomes the event's owner.
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        event  body      models.Event  true  "Event to create"
// @Success      201    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Security     CookieAuth
// @Router       /events [post]
func createEvent(c *gin.Context){

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data"})
		return
	}

	//event.ID=1
	userId := c.GetInt64("userId")
	event.UserID=userId

	err=event.Save()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"Event Created!", "event":event})
}

// updateEvent godoc
// @Summary      Update an event
// @Description  Update an event you own
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        id     path      int           true  "Event ID"
// @Param        event  body      models.Event  true  "Updated event fields"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string  "not the event owner"
// @Failure      500    {object}  map[string]string
// @Security     CookieAuth
// @Router       /events/{id} [put]
func updateEvent(c *gin.Context){
	eventId, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event."})
		return
	}
	if event.UserID != userId{
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return
	}
	updatedEvent.ID=eventId
	err=updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Event updated successfully!"})
}

// deleteEvent godoc
// @Summary      Delete an event
// @Description  Delete an event you own
// @Tags         events
// @Produce      json
// @Param        id   path      int  true  "Event ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string  "not the event owner"
// @Failure      500  {object}  map[string]string
// @Security     CookieAuth
// @Router       /events/{id} [delete]
func deleteEvent(c *gin.Context){
	eventId, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := c.GetInt64("userId")
	event, err:=models.GetEventByID(eventId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event."})
		return
	}
	if event.UserID != userId{
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete the event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Event deleted successfully!"})

}
/*
authenticated.GET("/myevent", myEvent)
func myEvent(c *gin.Context){
	//userId, err:= strconv.ParseInt.(c.Param("id"))
	//userId := c.GetInt64("userId")
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

}*/