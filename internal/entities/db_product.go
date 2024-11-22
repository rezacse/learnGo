package entities

type Product struct {
	ProductID string  `gorm:"primaryKey" json:"productId"`
	Name      string  `json:"name"`
	Price     float32 `germ:"type:numeric" json:"price"`
	VendorID  string  `json:"vendorId"`
}