package main

import (
	"os"

	"github.com/Carterc7/PortfolioApp/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Use the ShowHomePage function from main-controller.go
	router.GET("/", controllers.ShowHomePage)
	router.GET("/about", controllers.ShowAboutPage)
	router.GET("/projects", controllers.ShowProjectsPage)
	router.GET("/fantasy-about", controllers.ShowAboutFantasyPage)
	router.GET("/poll-about", controllers.ShowAboutPollPage)
	router.Static("/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // fallback for local development
	}

	router.Run(":" + port)
}
