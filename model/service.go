package model

// Service model
type Service struct {
	UUID          string     `gorm:"primaryKey" json:"uuid"`
	CustomerNIT   string     `gorm:"not null" json:"-"`
	TechnicianNIT string     `gorm:"not null" json:"-"`
	Customer      Customer   `gorm:"foreignKey:CustomerNIT" json:"customer"`
	Technician    Technician `gorm:"foreignKey:TechnicianNIT" json:"technician"`
	Comment       string     `json:"comment"`
	State         string     `gorm:"not null" json:"state"`
	Rate          int32      `gorm:"not null" json:"rate"`
}
