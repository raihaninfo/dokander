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
	r.LoadHTMLGlob("views/*.gohtml")

	r.NoRoute(a.notFound)
	r.GET("/ping", a.ping)
	r.GET("/", a.homeHandler)
	r.GET("/add-product", a.addProductHandler)
	r.POST("/add-product", a.addProductPostHandler)
	r.GET("/products", a.productsHandler)
	r.GET("/stook-out", a.stookOutHandler)
	r.GET("/add-customer", a.addCustomerHandler)

	r.Run(fmt.Sprintf(":%v", a.Server.Port)) // listen and serve on
	return r
}
