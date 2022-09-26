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

type Salary struct {
	Id         int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId     string
	EmployName string
	Month      string
	Amount     string
	PayDate    string
}

type ShopRent struct {
	Id      int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId  string
	Month   string
	Amount  string
	PayDate string
}

type UtilityBill struct {
	Id       int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId   string
	Note     string
	BillType string `gorm:"nullable"`
	BillId   string `gorm:"nullable"`
	Month    string
	Amount   string
	PayDate  string
}

type Entertainment struct {
	Id      int `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId  string
	Purpose string
	Amount  string
	Date    string
	Time    string
}

type Sales struct {
	Id          int `gorm:"primaryKey;autoIncrement:true;unique"`
	CustomersId int
	Product     string
	PaidAmount  float64
	DueAmount   float64
}
