package models

type Customers struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true;unique" json:"id"`
	ShopId      string `json:"shop_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	DateOfBirth string `json:"date_of_birth"`
	Reference   string `json:"reference"`
	Address     string `json:"address"`
	Status      int    `json:"status"`
}

type Products struct {
	Id              int    `gorm:"primaryKey;autoIncrement:true;unique" json:"id"`
	ShopId          int    `json:"shop_id"`
	PurchasesDate   string `json:"purchases_date"`
	ProductName     string `json:"product_name"`
	Brand           string `gorm:"nullable" json:"brand"`
	Model           string `gorm:"nullable" json:"model"`
	ProductType     string `gorm:"nullable" json:"product_type"`
	Quantity        int    `json:"quantity"`
	PurchasesAmount string `json:"purchases_amount"`
	SellAmount      string `json:"sell_amount"`
	Status          int    `json:"status"`
}

type Salary struct {
	Id         int    `gorm:"primaryKey;autoIncrement:true;unique" json:"id"`
	ShopId     string `json:"shop_id"`
	EmployName string `json:"employ_name"`
	Month      string `json:"month"`
	Amount     string `json:"amount"`
	PayDate    string `json:"pay_date"`
}

type ShopRent struct {
	Id      int    `gorm:"primaryKey;autoIncrement:true;unique"`
	ShopId  string `json:"shop_id"`
	Month   string `json:"month"`
	Amount  string `json:"amount"`
	PayDate string `json:"pay_date"`
}

type UtilityBill struct {
	Id       int    `gorm:"primaryKey;autoIncrement:true;unique" json:"id"`
	ShopId   string `json:"shop_id"`
	Note     string `json:"note"`
	BillType string `gorm:"nullable" json:"bill_type"`
	BillId   string `gorm:"nullable" json:"bill_id"`
	Month    string `json:"month"`
	Amount   string `json:"amount"`
	PayDate  string `json:"pay_date"`
}

type Entertainment struct {
	Id      int    `gorm:"primaryKey;autoIncrement:true;unique" json:"id"`
	ShopId  string `json:"shop_id"`
	Purpose string `json:"purpose"`
	Amount  string `json:"amount"`
	Date    string `json:"date"`
	Time    string `json:"time"`
}

type Sales struct {
	Id          int     `gorm:"primaryKey;autoIncrement:true;unique" json:"id"`
	CustomersId int     `json:"customers_id"`
	Product     string  `json:"product"`
	PaidAmount  float64 `json:"paid_amount"`
	DueAmount   float64 `json:"due_amount"`
}
