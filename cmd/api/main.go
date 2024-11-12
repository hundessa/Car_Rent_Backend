package main

import (
	"Car_Rent_Backend/internal/routes"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"

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
