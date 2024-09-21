package main

import (
	"Car_Rent_Backend/internal/routes"
	// "Car_Rent_Backend/internal/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// server := server.NewServer()
	r := gin.Default()

    routes.Routes(r) 

	// err := server.ListenAndServe()
	if err := r.Run(":8080"); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
