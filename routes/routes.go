package routes

import (
	"backend/handler"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	{
		auth.POST("/signup", handler.Signup)
		auth.POST("/login", handler.Login)
	}
	protected := r.Group("/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", handler.GetProfile)

		protected.GET("/:username", handler.GetUserByUsername)

		protected.POST("/matches", handler.CreateMatch)

		protected.POST("/teams", handler.CreateTeam)

		protected.POST("/teams/:id/players", handler.AddPlayersToTeam)
	}
}
