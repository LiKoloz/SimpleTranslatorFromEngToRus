package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Color struct {
	Eng string `json: "eng" binding:"required"`
	Rus string `json: "rus" binding:"required"`
}

var mapa map[string]string = map[string]string{
	"yellow": "желтый",
	"green":  "зеленый",
	"red":    "красный",
}

func main() {
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

	r.DELETE("/color/:color_name", func(c *gin.Context) {
		color_name := c.Param("color_name")
		delete(mapa, color_name)
		c.JSON(200, gin.H{
			"message": "Seccess delete!",
		})
	})

	r.POST("/color", func(c *gin.Context) {
		var color Color
		c.BindJSON(&color)
		mapa[color.Eng] = color.Rus
		c.JSON(201, gin.H{
			"message": "Succesfull create!",
		})
	})
	r.Run(":8080")
}
