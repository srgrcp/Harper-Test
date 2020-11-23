package model

// User base model for Customer and Technician
type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Nit  string
}

// Customer model
type Customer struct {
	User
}

// Technician model
type Technician struct {
	User
}
