package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/dokander/models"
)

func (a *application) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": a.AppName,
	})
}

func (a *application) notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.gohtml", nil)
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

func (a *application) addProductPostHandler(c *gin.Context) {
	shopId := 5
	name := c.Request.FormValue("name")
	brand := c.Request.FormValue("brand")
	model := c.Request.FormValue("model")
	productType := c.Request.FormValue("type")
	quantity := c.Request.FormValue("quantity")
	purchasePrice := c.Request.FormValue("purchases_price")
	sellPrice := c.Request.FormValue("sell_price")
	purchaseDate := c.Request.FormValue("purchases_date")

	quantityInt, err := a.toInt(quantity)
	if err != nil {
		log.Fatal(err)
	}

	products := models.Products{PurchasesDate: purchaseDate, ShopId: shopId, ProductName: name, Brand: brand, Model: model, ProductType: productType, Quantity: quantityInt, PurchasesAmount: purchasePrice, SellAmount: sellPrice}
	c.Bind(products)

	a.db.Create(&products)
	c.JSON(http.StatusOK, products)
	c.Header("Location", "/cmd")
	c.Redirect(http.StatusSeeOther, "/products")
}

func (a *application) productsHandler(c *gin.Context) {
	products := []models.Products{}
	a.db.Find(&products)
	c.HTML(http.StatusOK, "productList.gohtml", products)
}

func (a *application) stookOutHandler(c *gin.Context) {
	products := []models.Products{}
	a.db.Find(&products)
	c.HTML(http.StatusOK, "stookOut.gohtml", products)
}

func (a *application) addCustomerHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "addCustomer.gohtml", gin.H{
		"title": "Add Customer",
	})
}
