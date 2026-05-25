package routes

import (
	"time"

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
			"Accept",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
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

		// matches routes
		protected.POST("/matches", handler.CreateMatch)
		protected.POST("/matches/:id/toss", handler.StartMatchToss)

		// innings routes
		protected.POST("/matches/:id/innings", handler.StartInning)

		// teams routes
		protected.POST("/teams", handler.CreateTeams)
		protected.GET("/teams/:id/players", handler.GetTeamPlayers)
		protected.POST("/teams/:id/player", handler.AddPlayerToTeam)

		// protected user routes
		protected.GET("/profile/:username", handler.GetUserByUsername)

		// deliveries
		protected.POST("/innings/:inning_id/deliveries", handler.CreateDelivery)

	}
}
