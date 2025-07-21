package routes

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/controller"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/middleware"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, userController controller.UserController) {
	routes := r.Group("/api/user")
	{
		routes.POST("/register", userController.Register)
		routes.POST("/login", middleware.Authenticate(), userController.Login)
	}
}
