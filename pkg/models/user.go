package models

import (
	"github.com/androide18/go-user-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}
