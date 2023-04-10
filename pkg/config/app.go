package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "host=192.168.0.8 user=postgres password=postgres dbname=postgres port=5432"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
