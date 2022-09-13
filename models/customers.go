package models

type Customers struct {
	Id          int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId      string
	Name        string
	Email       string
	Mobile      string
	DateOfBirth string
	Reference   string
	Address     string
	Status      int
}
