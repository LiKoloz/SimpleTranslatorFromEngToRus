package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var mapa map[string]string = map[string]string{
	"yellow": "желтый",
	"green":  "зеленый",
	"red":    "красный",
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.GET("/color/:color-name", func(c *gin.Context) {
		color_name := c.Param("color-name")

		if val, ok := mapa[color_name]; ok {
			c.JSON(http.StatusOK, gin.H{
				"translation": val,
			})
		} else {
			c.JSON(http.StatusNotExtended, gin.H{
				"message": "color not found!",
			})
		}
	})
	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
