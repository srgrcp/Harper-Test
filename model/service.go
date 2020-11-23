package model

// Service model
type Service struct {
	UUID   string `gorm:"primaryKey"`
	UserID uint
	state  string
}
