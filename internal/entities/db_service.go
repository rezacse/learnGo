package entities

type Service struct {
	ServiceID string  `gorm:"primaryKey" json:"serviceId"`
	Name      string  `json:"name"`
	Price     float32 `germ:"type:numeric" json:"price"`
}