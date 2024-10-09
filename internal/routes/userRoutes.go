package routes

import (
    "github.com/gin-gonic/gin"
	"Car_Rent_Backend/internal/controllers"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/sign-up", controllers.SignUpHandler)
}