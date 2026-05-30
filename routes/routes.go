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
			"http://192.168.0.176:5173",
			"http://192.168.1.7:5173",
			"http://localhost:5173",
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
		MaxAge:           7 * 24 * time.Hour,
	}))
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handler.Signup)
		auth.POST("/login", handler.Login)
		auth.POST("/reset-password", handler.ResetPassword)
	}

	//users := r.Group("/users")
	//{
	//	users.GET("/players", handler.GetAllUsers)
	//}

	public := r.Group("/users")
	{
		public.GET("/players", handler.GetUsers) // getusers
		public.GET("/teams/:id/players", handler.GetTeamPlayers)
		public.GET("/profile/:username", handler.GetUserByUsername)
		public.GET("/teams", handler.GetTeams)

		//public.GET("/matches/:id/live", handler.GetLiveMatchState)

		// GetMatches / Live Scores endpoints
		// public.GET("/matches", handler.GetLiveMatches)
		// public.GET("/matches/:id", handler.GetMatch)
		public.GET("/matches/:id/live", handler.GetLiveMatchState)
		public.GET("/matches", handler.GetMatches)

	}

	protected := r.Group("/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", handler.GetProfile)

		// matches routes
		protected.POST("/matches", handler.CreateMatch)
		protected.POST("/matches/:id/toss", handler.StartMatchToss)
		protected.POST("/matches/start", handler.StartMatch)
		protected.POST("/matches/setup", handler.SuperSetupMatchHandler)

		// innings routes
		protected.POST("/matches/:id/innings", handler.StartInning)

		// teams routes
		protected.POST("/teams", handler.CreateTeams)
		//protected.GET("/teams", handler.GetTeam)
		//protected.GET("/teams/:id/players", handler.GetTeamPlayers)
		protected.POST("/teams/:id/player", handler.AddPlayerToTeam)

		// user routes
		//protected.GET("/profile/:username", handler.GetUserByUsername)

		// deliveries
		//protected.POST("/deliveries", handler.CreateDelivery)
		protected.POST("/innings/:inning_id/deliveries", handler.CreateDelivery)

	}
}
