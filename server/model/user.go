package model

type User struct {
	ID 	string `gorm:"unique"`
	Username 	string `gorm:"unique"`
	Password string `gorm:"not null"`
}