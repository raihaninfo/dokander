package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func (a *application) router() http.Handler {
	r := gin.Default()
	r.Use(static.Serve("/assets", static.LocalFile("./public", true)))
	r.LoadHTMLGlob("views/*.html")

	r.GET("/ping", a.ping)
	r.GET("/", a.homeHandler)

	r.Run(fmt.Sprintf(":%v", a.Server.Port)) // listen and serve on
	return r
}
