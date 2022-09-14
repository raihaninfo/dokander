package models

type ShopRent struct {
	Id      int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId  string
	Month   string
	Amount  string
	PayDate string
}
