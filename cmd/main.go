package main

import (
	"backend/handler"
	"log"

	"backend/database"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	database.ConnectDB()

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/matches", handler.CreateMatch)

	r.POST("/teams", handler.CreateTeam)
	routes.SetupRoutes(r)

	for _, route := range r.Routes() {
		log.Println(route.Method, route.Path)
	}

	r.Run(":8080")
}
