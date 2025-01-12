package model

import "time"

type Book struct {
	ID            string     `gorm:"type:char(36);primaryKey;not null"`
	Title         string     `gorm:"type:varchar(255);not null"`
	Author        string     `gorm:"type:varchar(255);not null"`
	PublishedDate time.Time  `gorm:"not null"`
	Status        string     `gorm:"type:varchar(255);default:'Available';not null"`
	UserID        *string    `gorm:"type:char(36);default:null"`
}

type BorrowedBook struct {
	ID           string     `gorm:"type:char(36);primaryKey;not null"`
	BookID       string     `gorm:"type:char(36);not null;constraint:OnDelete:CASCADE"`
	UserID       string     `gorm:"type:char(36);not null;constraint:OnDelete:CASCADE"`
	BorrowedDate time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	ReturnDate   *time.Time `gorm:"default:null"`
}