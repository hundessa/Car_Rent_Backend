package controllers

import (
	"Car_Rent_Backend/internal/migrations"
	"Car_Rent_Backend/internal/models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)


func CarCreateHandler(c *gin.Context) {

	var cars models.Cars

	if err := c.ShouldBindBodyWithJSON(&cars); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("error", err.Error())
		return
	}

	if err := cars.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("error", err.Error())
		return
	}

	db := migrations.ConnectDB()
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		log.Println("Database connection error")
		return
	}

	if err := db.Create(&cars).Error; err != nil {

		if strings.Contains(err.Error(), "duplicate key violates unique constraint") {
			c.JSON(400, gin.H{"error": "Car Model already exists"})
			log.Println("Car Model already exists")
			return
		}

		c.JSON(500, gin.H{"error": "Internal server error"})
		log.Println("Internal server error")
		return
	}

	c.JSON(200, gin.H{"message": "Car creation successful"})
	log.Println("Car creation successful")
	
}