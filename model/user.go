package model

// User base model for Customer and Technician
type User struct {
	Nit  string `gorm:"primaryKey" json:"nit"`
	Name string `gorm:"not null" json:"name"`
}

// Customer model
type Customer struct {
	User
	Service []Service `gorm:"foreignKey:CustomerNIT" json:",omitempty"`
}

// Technician model
type Technician struct {
	User
	Service []Service `gorm:"foreignKey:TechnicianNIT" json:",omitempty"`
}
