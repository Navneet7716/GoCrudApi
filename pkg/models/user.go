package models

import (
	"github.com/navneet7716/go-bookstore/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     uint   `json:"phone"`
}

func init() {

	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})

}

func (u *User) CreateUser() *User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	u.Password = string(hashed)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id = ?", Id).Find(&getUser)

	return &getUser, db

}
