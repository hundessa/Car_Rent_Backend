package controllers

import (
	"Car_Rent_Backend/internal/middlewares"
	"Car_Rent_Backend/internal/migrations"
	"Car_Rent_Backend/internal/models"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage" // Import Firebase storage
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// UploadImageToFirebase uploads an image to Firebase Storage
func UploadImageToFirebase(fileHeader *multipart.FileHeader, fileName string) (string, error) {
    // Initialize Firebase storage client
    storageClient, err := firebase.InitFirebase() // Assume you have Firebase initialized in this package
    if err != nil {
        return "", err
    }

    // Firebase credentials file (JSON key file path)
    // credsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
    credsFile := os.Getenv("/home/hundessa/Car_Rent_Backend/internal/config/car-rent-d1c84-firebase-adminsdk-2h0ti-fac4fab7d6.json")
    fmt.Println("Using GOOGLE_APPLICATION_CREDENTIALS:", credsFile)
    if credsFile == "" {
        return "", fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS is not set")
    }

    // Initialize Firebase storage client
    ctx := context.Background()
    client, err := storage.NewClient(ctx, option.WithCredentialsFile(credsFile))
    if err != nil {
        return "", err
    }
    defer client.Close()

    file, err := fileHeader.Open()
    if err != nil {
        return "", err
    }
    defer file.Close()

    bucket, err := storageClient.Bucket("gs://car-rent-d1c84.appspot.com")
    if err != nil {
        return "", err
    }

    object := bucket.Object(fmt.Sprintf("images/%s", fileName)).NewWriter(context.Background())
    if _, err := io.Copy(object, file); err != nil {
        return "", err
    }

    if err := object.Close(); err != nil {
        return "", err
    }

        const bucketName = "gs://car-rent-d1c84.appspot.com"

    // Return the public URL of the uploaded file
    url := fmt.Sprintf("https://storage.googleapis.com/%s/images/%s", bucketName, fileName)
    return url, nil
}

func CarCreateHandler(c *gin.Context) {
    var cars models.Cars

 // Extract form data fields (use c.PostForm for form-data)
 cars.CarName = c.PostForm("carname")
 cars.CarModel = c.PostForm("carmodel")
 cars.CarProductionYear = c.PostForm("carproductionyear")
 cars.CarMileage = c.PostForm("carmileage")
 cars.Description = c.PostForm("description")
 cars.CarPrice = c.PostForm("carprice")
 cars.CarRating = c.PostForm("carrating")

    // Bind the form data (expecting multipart form)
    if err := c.ShouldBind(&cars); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Validate the car data
    if err := cars.Validate(); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Get the image from the form data
    fileHeader, err := c.FormFile("carimage")
    if err != nil {
        c.JSON(400, gin.H{"error": "Image file is required"})
        return
    }

    // Create a unique file name for the image
    fileName := fmt.Sprintf("%s-%d%s", strings.ToLower(cars.CarModel), time.Now().Unix(), filepath.Ext(fileHeader.Filename))

    // Upload the image to Firebase Storage
    imageURL, err := UploadImageToFirebase(fileHeader, fileName)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to upload image"})
        return
    }

    // Set the image URL in the Car model
    cars.CarImage = imageURL

    // Connect to the database
    db := migrations.ConnectDB()
    if db == nil {
        c.JSON(500, gin.H{"error": "Database connection failed"})
        return
    }

    // Save the car details to the database
    if err := db.Create(&cars).Error; err != nil {
        if strings.Contains(err.Error(), "duplicate key violates unique constraint") {
            c.JSON(400, gin.H{"error": "Car Model already exists"})
            return
        }
        c.JSON(500, gin.H{"error": "Internal server error"})
        return
    }

    c.JSON(200, gin.H{"message": "Car creation successful", "car": cars})
}
