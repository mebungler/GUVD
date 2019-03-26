package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../models"
)

var DB *gorm.DB

func ConnectAndCreate(databaseType, connectionString string) {
	var err error
	DB, err = gorm.Open(databaseType, connectionString)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.Building{})
	DB.AutoMigrate(&models.Property{})
	DB.AutoMigrate(&models.Person{})
	DB.LogMode(true)
}

func Close() {
	DB.Close()
}
