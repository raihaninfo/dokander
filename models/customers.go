package models

type Customers struct {
	Id          int `gorm:"primaryKey;autoIncrement:true;unique"`
	Name        string
	Email       string `gorm:"unique"`
	Mobile      string
	DateOfBirth string
	Reference   string
	Address     string
	Status      int
}
