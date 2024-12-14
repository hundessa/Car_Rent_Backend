package main

import (
	"Car_Rent_Backend/internal/helpers"
	"Car_Rent_Backend/internal/models"
	"Car_Rent_Backend/internal/routes"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"Car_Rent_Backend/internal/migrations"
)

func main() {

	// Call the connectDB function
	db := migrations.ConnectDB() // Use the connectDB function from the migrations package

	if db != nil {
		fmt.Println("Database connected successfully!")
	} else {
		fmt.Println("Failed to connect to the database.")
	}

	// Check if an admin exists and create one if not
	checkAndCreateAdmin(db)

	// server := server.NewServer()
	r := gin.Default()

	// Configure CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your frontend's URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.UserRoutes(r)
	routes.CarRoutes(r)

	// err := server.ListenAndServe()
	if err := r.Run(":8080"); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}

// checkAndCreateAdmin checks if an admin exists; if not, it creates one
func checkAndCreateAdmin(db *gorm.DB) {
	var admin models.User
	err := db.Where("role = ?", "admin").First(&admin).Error

	// If no admin found, create a new one
	if err == gorm.ErrRecordNotFound {
		newAdmin := models.User{
			FirstName: "admin",
			Email:     "admin@mail.com",
			Password:  "admin@ca", // You should hash the password here
			Role:      "admin",
		}

		// Hash the password (replace with actual hashing)
		hashedPassword, err := helpers.HashAndSalt([]byte(newAdmin.Password))
		if err != nil {
			fmt.Println("Error hashing password:", err)
			return // Add this line to exit the function if there's an error hashing the password
		}
		newAdmin.Password = hashedPassword

		if err := db.Create(&newAdmin).Error; err != nil {
			fmt.Println("Error creating admin:", err)
		} else {
			fmt.Println("Admin user created successfully!")
		}
	} else if err != nil {
		fmt.Println("Error checking for admin:", err)
	} else {
		fmt.Println("Admin already exists.")
	}
}
