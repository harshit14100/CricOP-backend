package routes

import (
	"backend/handler"
	"backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
	}))
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handler.Signup)
		auth.POST("/login", handler.Login)
		auth.POST("/reset-password", handler.ResetPassword)
	}

	users := r.Group("/users")
	{
		users.GET("/players", handler.GetAllUsers)
	}
	protected := r.Group("/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", handler.GetProfile)

		//protected.GET("/players", handler.GetAllUsers)

		protected.POST("/matches", handler.CreateMatch)

		protected.POST("/teams", handler.CreateTeams)

		protected.GET("/teams/:id/players", handler.GetTeamPlayers)

		protected.POST("/teams/:id/player", handler.AddPlayerToTeam)

		protected.GET("/profile/:username", handler.GetUserByUsername)

	}
}
