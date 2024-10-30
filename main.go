package main

import (
	"event-go-rest-api/models" // import the models package
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	server := gin.Default()

	// server.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"message": "Hello, World!"})
	// })

	server.GET("/", getHelloWorld)
	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getHelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()

	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	fmt.Println("Creating event...")
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// event.CreateEvent()
	event.ID = 1
	event.UserID = 10

	event.CreateEvent()

	ctx.JSON(http.StatusCreated, event)
}
