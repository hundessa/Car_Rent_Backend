package migrations

import (
	"Car_Rent_Backend/internal/helpers"
	"Car_Rent_Backend/internal/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)


func ConnectDB() *gorm.DB {


	godotenv.Load("/home/hundessa/Car_Rent_Backend/.env")
	// godotenv.Load(".env")

	database   := os.Getenv("DB_DATABASE")
	password   := os.Getenv("DB_PASSWORD")
	username   := os.Getenv("DB_USERNAME")
	port       := os.Getenv("DB_PORT")
	host       := os.Getenv("DB_HOST")


 // Ensure the variables are not empty
 if host == "" || port == "" || username == "" || password == "" || database == "" {
	fmt.Println("Some environment variables are missing or empty.")
	return nil
}

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		host, port, username, password, database)

		db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		helpers.HandleError(err)
	}

	// AutoMigrate to create the User table
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Cars{})

	if err := db.AutoMigrate(&models.User{}, &models.Cars{}); err != nil {
		fmt.Println("Migration error: ", err)
	}
	
	return db
}