package orm

import (
	"github.com/srgrcp/Harper-Test/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initialize and migrate database
func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("db/sqlite3.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Technician{})
	db.AutoMigrate(&model.Service{})
}
