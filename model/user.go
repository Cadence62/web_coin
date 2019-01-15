package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName     string
	Email        string `gorm:"size:255;unique_index;not null"`
	Phone        string `gorm:"size:25;unique_index"`
	ApiKey       string `gorm:"size:50;unique_index"`
	ApiSecretKey string `gorm:"size:50;unique_index"`
}
