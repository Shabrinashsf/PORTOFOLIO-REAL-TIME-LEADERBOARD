package main

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/controller"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/initializers"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/repository"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/routes"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/service"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgre()
	initializers.NewRedisClient()
}

func main() {
	var (
		// Implementation Dependency Injection
		// Repository
		userRepository        repository.UserRepository        = repository.NewUserRepository(initializers.DB)
		leaderboardRepository repository.LeaderboardRepository = repository.NewLeaderboardRepository(initializers.NewRedisClient(), initializers.DB)

		// Service
		userService        service.UserService        = service.NewUserService(userRepository)
		leaderboardService service.LeaderboardService = service.NewLeaderboardService(leaderboardRepository)

		// Controller
		userController        controller.UserController        = controller.NewUserController(userService)
		leaderboardController controller.LeaderboardController = controller.NewLeaderboardController(leaderboardService)
	)

	r := gin.Default()
	routes.User(r, userController)
	routes.Leaderboard(r, leaderboardController)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
