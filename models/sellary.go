package models

type Salary struct {
	Id         int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId     string
	EmployName string
	Month      string
	Amount     string
	PayDate    string
}
