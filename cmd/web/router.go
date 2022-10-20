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
	r.GET("/", a.homeHandler)
	r.GET("/add-product", a.addProductHandler)
	r.POST("/add-product", a.addProductPostHandler)
	r.GET("/products", a.productsHandler)
	r.GET("/products/update/:id", a.productsUpdateHandler)
	r.POST("/products/update/:id", a.productsUpdateAuthHandler)
	r.GET("/stook-out", a.stookOutHandler)
	r.GET("/add-customer", a.addCustomerHandler)
	r.POST("/add-customer", a.addCustomerPostHandler)
	r.GET("/customers", a.customers)
	r.GET("/customer/update/:id", a.customerUpdate)
	r.POST("/customer/update/:id", a.customerPostUpdate)

	r.GET("/shop-rent", a.shopRent)
	r.GET("/add-shop-rent", a.addShopRent)
	r.POST("/add-shop-rent", a.addShopRentPost)

	r.GET("/salary", a.salary)
	r.GET("/add-salary", a.addSalary)
	r.POST("/add-salary", a.addSalaryPost)

	r.GET("/utility-bill", a.utilityBill)
	r.GET("/add-utility-bill", a.addUtilityBill)
	r.POST("/add-utility-bill", a.addUtilityBillPost)

	r.GET("/entertainment-bill", a.entertainmentBill)
	r.GET("/add-entertainment-bill", a.addEntertainmentBill)
	r.POST("/add-entertainment-bill", a.addEntertainmentBillPost)

	r.GET("/sales", a.productsSells)
	r.POST("/sales", a.productsSellsPost)

	// api
	api := r.Group("/api")
	{
		api.GET("/get-customers", a.getCustomers)
		api.GET("/get-products", a.getProducts)
	}

	r.Run(fmt.Sprintf(":%v", a.Server.Port)) // listen and serve on
	return r
}
