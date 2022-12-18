package models

import (
	"github.com/SujitSarkar/Go-Practice/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Boook struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json: "author"`
	Publication string `json: "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Boook{})
}
