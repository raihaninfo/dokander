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

func (a *application) addProductHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "addProduct.gohtml", gin.H{
		"title": "Add Product",
	})
}

func (a *application) productsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "productList.gohtml", gin.H{
		"title": "Products",
	})
}

func (a *application) stookOutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "stookOut.gohtml", gin.H{
		"title": "Stook Out",
	})
}

func (a *application) addCustomerHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "addCustomer.gohtml", gin.H{
		"title": "Add Customer",
	})
}
