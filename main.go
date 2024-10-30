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

	server.Run(":8080") // localhost:8080
}

func getHelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()

	ctx.JSON(http.StatusOK, events)
}
