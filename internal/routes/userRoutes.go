package routes

import (
	"Car_Rent_Backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// Group routes under /user for organization
	user := r.Group("/user")
	{
		user.POST("/sign-up", controllers.SignUpHandler)
		user.POST("/sign-in", controllers.SigninHandler)
	}
}
