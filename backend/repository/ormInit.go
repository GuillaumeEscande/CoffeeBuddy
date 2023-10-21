package repository


import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

	"coffeebuddy/model"
)

var DB

func InitOrm(){
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}


	// Migrate the schema
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.CoffeeMachine{})
	DB.AutoMigrate(&model.CoffeeCreditEvent{})
}