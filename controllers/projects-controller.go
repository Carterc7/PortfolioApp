package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowProjectsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "projects.html", gin.H{
		"title": "Projects | Carter Campbell",
	})
}
