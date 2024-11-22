package entities

type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"LastName"`
	Email      string `json:"emailAddress"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}