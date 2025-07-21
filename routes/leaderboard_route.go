package routes

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/controller"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/middleware"

	"github.com/gin-gonic/gin"
)

func Leaderboard(r *gin.Engine, leaderboardController controller.LeaderboardController) {
	routes := r.Group("/api")
	{
		routes.POST("/submit", middleware.Authenticate(), leaderboardController.SubmitScore)
		routes.PUT("/update", middleware.Authenticate(), leaderboardController.UpdateScore)
		routes.GET("/me/:game", middleware.Authenticate(), leaderboardController.GetRankByGame)
		routes.GET("/leaderboard/:game/:limit", middleware.Authenticate(), leaderboardController.GetLeaderboard)
	}
}
