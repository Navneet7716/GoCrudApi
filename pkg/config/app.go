package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
