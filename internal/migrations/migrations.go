package migrations

import (
	"Car_Rent_Backend/internal/helpers"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	FirstName string
	LastName string
	Email string
	Password string
}


func ConnectDB() *gorm.DB {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file:", err) // Print error if .env file is not loaded
		return nil // Return nil if the DB connection cannot be established
	} else {
        fmt.Println(".env file loaded successfully")
    }

	database   := os.Getenv("DB_DATABASE")
	password   := os.Getenv("DB_PASSWORD")
	username   := os.Getenv("DB_USERNAME")
	port       := os.Getenv("DB_PORT")
	host       := os.Getenv("DB_HOST")


	// Debugging output after loading environment variables
fmt.Printf("Host: %s\n", os.Getenv("DB_HOST"))
fmt.Printf("Port: %s\n", os.Getenv("DB_PORT"))
fmt.Printf("Username: %s\n", os.Getenv("DB_USERNAME"))
fmt.Printf("Password: %s\n", os.Getenv("DB_PASSWORD"))
fmt.Printf("Database: %s\n", os.Getenv("DB_DATABASE"))

 // Ensure the variables are not empty
 if host == "" || port == "" || username == "" || password == "" || database == "" {
	fmt.Println("Some environment variables are missing or empty.")
	return nil
}

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
		db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		helpers.HandleError(err)
	}
	return db
}