package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *application) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": a.AppName,
	})
}

func (a *application) homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title": "Home Page",
	})
}
