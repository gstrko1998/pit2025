package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(100)"`
	LastName    string `gorm:"type:varchar(100)"`
	Address     string `gorm:"type:varchar(200)"`
	City        string `gorm:"type:varchar(100)"`
	State       string `gorm:"type:varchar(100)"`
	PostalCode  string `gorm:"type:varchar(20)"`
	PhoneNumber string `gorm:"type:varchar(20)"`
	UserID      uint
	User        User // Belongs to relationship with User
}

func (p *Profile) FullName() string {
	return p.FirstName + " " + p.LastName
}
