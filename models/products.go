package models

type Products struct {
	Id              int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId          int
	PurchasesDate   string
	ProductName     string
	Brand           string `gorm:"nullable"`
	Model           string `gorm:"nullable"`
	ProductType     string `gorm:"nullable"`
	Quantity        int
	PurchasesAmount string
	SellAmount      string
	Status          int
}
