package main

import (
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
	router.Static("/assets", "./assets")

	router.Run(":8081") // Runs on localhost:8081
}
