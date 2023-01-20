package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	GivenName string `gorm:"not null" json:"given_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null; unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
