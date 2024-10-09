package routes

import (
	"github.com/gin-gonic/gin"
	"Car_Rent_Backend/internal/controllers"
)

func CarRoutes(r *gin.Engine) {
	r.POST("/car-create", controllers.CarCreateHandler)
}