package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"title": "About | Carter Campbell",
	})
}

func ShowAboutFantasyPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about-fantasy.html", gin.H{
		"title": "Fantasy | About",
	})
}
