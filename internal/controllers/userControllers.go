package controllers

import (
	"Car_Rent_Backend/internal/migrations"
	"Car_Rent_Backend/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {

	var user models.User

    // Bind JSON to the User struct
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }   

    // Validate the user struct
    if err := user.Validate(); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Get the DB connection
	db := migrations.ConnectDB()
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}

	// Save the user to the database
	if err := db.Create(&user).Error; err != nil {
		
        // Handle unique constraint violation error (duplicate email)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(400, gin.H{"error": "Email already exists"})
			return
		}
		// Handle other possible errors
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
        
        // c.JSON(500, gin.H{"error": "Failed to save user to database"})
		// return
	}

	// Respond with a success message
	c.JSON(200, gin.H{"message": "Sign up successful"})

}