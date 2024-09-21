package controllers

import (
	"github.com/gin-gonic/gin"
"Car_Rent_Backend/internal/models"
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

    // Process the user registration (e.g., save to database)
    println("User signed up:", user.FirstName, user.LastName)

    // Respond with a success message
    c.JSON(200, gin.H{"message": "Sign up successful"})

}