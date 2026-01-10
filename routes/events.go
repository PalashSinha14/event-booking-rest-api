package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/models"
)

func getEvents(c *gin.Context) {
	events,err:=models.GetAllEvents()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	c.JSON(http.StatusOK, events)
}

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

func createEvent(c *gin.Context){
	var event models.Event
	err:=c.ShouldBindJSON(&event)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data"})
		return
	}
	event.ID=1
	event.UserID=1
	err=event.Save()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"Event Created!", "event":event})
}

func updateEvent(c *gin.Context){
	eventId, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	_, err=models.GetEventByID(eventId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event."})
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

func deleteEvent(c *gin.Context){
	eventId, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err:=models.GetEventByID(eventId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event."})
		return
	}
	err = event.Delete()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete the event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Event deleted successfully!"})

}

