package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/dokander/models"
)

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
	shopId := 8
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
	c.Redirect(http.StatusSeeOther, "/products")
}

func (a *application) productsHandler(c *gin.Context) {
	products := []models.Products{}
	a.db.Find(&products)
	c.HTML(http.StatusOK, "productList.gohtml", products)
}

func (a *application) productsUpdateHandler(c *gin.Context) {
	product := models.Products{}
	id := c.Param("id")
	a.db.Find(&product, id)
	c.HTML(http.StatusOK, "updateProduct.gohtml", product)
}

func (a *application) productsUpdateAuthHandler(c *gin.Context) {
	productId := c.Request.FormValue("id")
	name := c.Request.FormValue("name")
	brand := c.Request.FormValue("brand")
	model := c.Request.FormValue("model")
	productType := c.Request.FormValue("type")
	quantity := c.Request.FormValue("quantity")
	purchasePrice := c.Request.FormValue("purchases_price")
	sellPrice := c.Request.FormValue("sell_price")
	purchaseDate := c.Request.FormValue("purchases_date")
	fmt.Println(productId, name, brand, model, productType, quantity, purchasePrice, sellPrice, purchaseDate)

	id := c.Param("id")
	var products models.Products
	if err := a.db.Where("id = ?", id).First(&products).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	quantityInt, err := a.toInt(quantity)
	if err != nil {
		log.Fatal(err)
	}
	products.ProductName = name
	products.Brand = brand
	products.Model = model
	products.ProductType = productType
	products.Quantity = quantityInt
	products.PurchasesAmount = purchasePrice
	products.SellAmount = sellPrice
	products.PurchasesDate = purchaseDate
	c.Bind(&products)
	a.db.Save(&products)
	c.Redirect(http.StatusSeeOther, "/products")

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
