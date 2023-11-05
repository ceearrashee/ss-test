package models

import (
	"gorm.io/gorm"
)

type (
	// User struct represents the user model in the database.
	User struct {
		gorm.Model
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}
)
