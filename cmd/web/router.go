package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *application) router() http.Handler {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on
	return r
}
