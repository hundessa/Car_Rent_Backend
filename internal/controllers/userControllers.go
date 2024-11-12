// package controllers

// import (
// 	"Car_Rent_Backend/internal/helpers"
// 	"Car_Rent_Backend/internal/migrations"
// 	"Car_Rent_Backend/internal/models"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func SignUpHandler(c *gin.Context) {

// 	var user models.User

// 	// Bind JSON to the User struct
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Validate the user struct
// 	if err := user.Validate(); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Hash the password
// 	user.Password = helpers.HashAndSalt([]byte(user.Password))

// 	// Get the DB connection
// 	db := migrations.ConnectDB()
// 	if db == nil {
// 		c.JSON(500, gin.H{"error": "Database connection failed"})
// 		return
// 	}

// 	// Save the user to the database
// 	if err := db.Create(&user).Error; err != nil {
// 		helpers.HandleError(err)

// 		// Handle unique constraint violation error (duplicate email)
// 		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
// 			c.JSON(400, gin.H{"message": "Email already exists"})
// 			fmt.Print("Email already exists")
// 			return
// 		}
// 		// Handle other possible errors
// 		c.JSON(500, gin.H{"error": "Internal server error"})
// 		return

// 		// c.JSON(500, gin.H{"error": "Failed to save user to database"})
// 		// return
// 	}

// 	// Respond with a success message
// 	c.JSON(200, gin.H{"message": "Sign up successful"})
// 	print("Sign up successful")

// }

// func SigninHandler(c *gin.Context) {
// 	var loginData struct {
// 		Email    string `json:"email" binding:"required,email"`
// 		Password string `json:"password" binding:"required"`
// 	}

// 	// Bind JSON from request body to loginData struct
// 	if err := c.ShouldBindJSON(&loginData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
// 		fmt.Println("Error binding JSON:", err) // Debug log
// 		return
// 	}

// 	fmt.Println("Email:", loginData.Email)       // Debug log
// 	fmt.Println("Password:", loginData.Password) // Debug log

// 	// Get the DB connection
// 	db := migrations.ConnectDB()
// 	if db == nil {
// 		fmt.Println("Database connection failed") // Debug log
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database connection failed"})
// 		return
// 	}

// 	// Find the user by email
// 	var user models.User
// 	if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
// 		fmt.Println("Error finding user:", err) // Debug log
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
// 		}
// 		return
// 	}

// 	// Validate the password
// 	if !helpers.CheckPasswordHash(loginData.Password, user.Password) {
// 		fmt.Println("Invalid credentials") // Debug log
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Signin successful"})
// }

package controllers

import (
	"Car_Rent_Backend/internal/helpers"
	"Car_Rent_Backend/internal/migrations"
	"Car_Rent_Backend/internal/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUpHandler(c *gin.Context) {
	var user models.User

	// Bind JSON to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the user struct
	if err := user.Validate(); err != nil {
		log.Printf("Validation error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	// user.Password = helpers.HashAndSalt([]byte(user.Password))
	hashedPassword, err := helpers.HashAndSalt([]byte(user.Password))
	if err != nil {
		log.Printf("Password hashing error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	// Get the DB connection
	db := migrations.ConnectDB()
	if db == nil {
		log.Println("Database connection failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	// Save the user to the database
	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error saving user to database: %v\n", err)
		helpers.HandleError(err)

		// Handle unique constraint violation (duplicate email)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Println("Email already exists")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
			log.Println("Response sent: Email already exists")
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Log and respond with a success message
	log.Println("Sign up successful")
	c.JSON(http.StatusOK, gin.H{"message": "Sign up successful"})
}

func SigninHandler(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON from request body to loginData struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		fmt.Println("Error binding JSON:", err) // Debug log
		return
	}

	fmt.Println("Email:", loginData.Email)       // Debug log
	fmt.Println("Password:", loginData.Password) // Debug log

	// Get the DB connection
	db := migrations.ConnectDB()
	if db == nil {
		fmt.Println("Database connection failed") // Debug log
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database connection failed"})
		return
	}

	// Find the user by email
	var user models.User
	if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		fmt.Println("Error finding user:", err) // Debug log
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}
		return
	}

	// Validate the password
	if !helpers.CheckPasswordHash(loginData.Password, user.Password) {
		fmt.Println("Invalid credentials") // Debug log
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signin successful"})
	fmt.Println("Signin successful")
}
