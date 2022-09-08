package models

type Products struct {
	Id              int `gorm:"primaryKey;autoIncrement:true;unique"`
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
