package main

import (
	"Car_Rent_Backend/internal/routes"
	// "Car_Rent_Backend/internal/server"
	"fmt"

	"github.com/gin-gonic/gin"
	_"github.com/lib/pq"
    _"github.com/joho/godotenv"

	"Car_Rent_Backend/internal/migrations"
)

func main() {

	// Call the connectDB function
    db := migrations.ConnectDB()  // Use the connectDB function from the migrations package

    if db != nil {
        fmt.Println("Database connected successfully!")
    } else {
        fmt.Println("Failed to connect to the database.")
    }
	// server := server.NewServer()
	r := gin.Default()

    routes.UserRoutes(r) 
	routes.CarRoutes(r)

	// err := server.ListenAndServe()
	if err := r.Run(":8080"); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
