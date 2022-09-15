package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/dokander/models"
)

func (a *application) notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.gohtml", nil)
}

func (a *application) homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", nil)
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
	name := c.Request.FormValue("name")
	brand := c.Request.FormValue("brand")
	model := c.Request.FormValue("model")
	productType := c.Request.FormValue("type")
	quantity := c.Request.FormValue("quantity")
	purchasePrice := c.Request.FormValue("purchases_price")
	sellPrice := c.Request.FormValue("sell_price")
	purchaseDate := c.Request.FormValue("purchases_date")

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

func (a *application) addCustomerPostHandler(c *gin.Context) {
	shopId := "1"
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	mobile := c.Request.FormValue("mobile")
	dateOfBirth := c.Request.FormValue("date")
	reference := c.Request.FormValue("reference")
	address := c.Request.FormValue("address")
	customer := models.Customers{ShopId: shopId, Name: name, Email: email, Mobile: mobile, DateOfBirth: dateOfBirth, Reference: reference, Address: address}

	c.Bind(customer)
	a.db.Create(&customer)
	c.Redirect(http.StatusSeeOther, "/customers")

}

func (a *application) customers(c *gin.Context) {
	customers := []models.Customers{}
	a.db.Find(&customers)
	c.HTML(http.StatusOK, "customersList.gohtml", customers)
}

func (a *application) customerUpdate(c *gin.Context) {
	id := c.Param("id")
	customer := models.Customers{}
	a.db.Find(&customer, id)
	c.HTML(http.StatusOK, "updateCustomer.gohtml", customer)
}

func (a *application) customerPostUpdate(c *gin.Context) {
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	mobile := c.Request.FormValue("mobile")
	dateOfBirth := c.Request.FormValue("date")
	reference := c.Request.FormValue("reference")
	address := c.Request.FormValue("address")

	id := c.Param("id")
	var customers models.Customers
	if err := a.db.Where("id = ?", id).First(&customers).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	customers.Name = name
	customers.Email = email
	customers.Mobile = mobile
	customers.DateOfBirth = dateOfBirth
	customers.Reference = reference
	customers.Address = address

	c.Bind(&customers)
	a.db.Save(&customers)
	c.Redirect(http.StatusSeeOther, "/customers")
}

func (a *application) shopRent(c *gin.Context) {
	rent := []models.ShopRent{}
	a.db.Find(&rent)
	c.HTML(http.StatusOK, "shopRent.gohtml", rent)
}

func (a *application) addShopRent(c *gin.Context) {
	c.HTML(http.StatusOK, "addShopRent.gohtml", gin.H{})
}

func (a *application) addShopRentPost(c *gin.Context) {
	shopId := "1"
	month := c.Request.FormValue("month")
	amount := c.Request.FormValue("amount")
	date := c.Request.FormValue("date")

	rent := models.ShopRent{ShopId: shopId, Month: month, Amount: amount, PayDate: date}
	c.Bind(&rent)
	a.db.Create(&rent)
	c.Redirect(http.StatusSeeOther, "/shop-rent")

}

func (a *application) salary(c *gin.Context) {
	salary := []models.Salary{}
	a.db.Find(&salary)
	c.HTML(http.StatusOK, "salary.gohtml", salary)
}

func (a *application) addSalary(c *gin.Context) {
	c.HTML(http.StatusOK, "addSalary.gohtml", gin.H{})
}

func (a *application) addSalaryPost(c *gin.Context) {
	shopId := "1"
	name := c.Request.FormValue("name")
	month := c.Request.FormValue("month")
	amount := c.Request.FormValue("amount")
	date := c.Request.FormValue("date")

	salary := models.Salary{ShopId: shopId, EmployName: name, Month: month, Amount: amount, PayDate: date}
	c.Bind(&salary)
	err := a.db.Create(&salary).Error
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusSeeOther, "/salary")

}

func (a *application) utilityBill(c *gin.Context) {
	bill:= []models.UtilityBill{}
	a.db.Find(&bill)
	c.HTML(http.StatusOK, "utilityBill.gohtml", bill)
}

func (a *application) addUtilityBill(c *gin.Context) {
	c.HTML(http.StatusOK, "addUtilityBill.gohtml", gin.H{})
}

func (a *application) addUtilityBillPost(c *gin.Context) {
	shopId:= "1"
	note:= c.Request.FormValue("note")
	billType:= c.Request.FormValue("bill-type")
	billId:= c.Request.FormValue("bill-id")
	month:= c.Request.FormValue("month")
	amount:= c.Request.FormValue("amount")
	payDate:= c.Request.FormValue("pay-date")

	bill:= models.UtilityBill{ShopId: shopId, Note: note, BillType: billType, BillId: billId, Month: month, Amount: amount, PayDate: payDate}

	c.Bind(&bill)
	err:= a.db.Create(&bill).Error
	if err!=nil{
		log.Fatal(err)
	}
	c.Redirect(http.StatusSeeOther, "/utility-bill")
}
