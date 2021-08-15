package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Task{}, &User{}, &Status{}, &History{})

	// fill in the status table if it is empty
	var values = [5]string{"New", "Active", "Stopped", "Done", "Delete"}
	for index, value := range values {
		var status = Status{ID: uint(index + 1), Name: value}
		database.FirstOrCreate(&status)
	}

	DB = database
}
