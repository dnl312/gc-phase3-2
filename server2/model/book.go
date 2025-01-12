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
	Bookid       string     `gorm:"type:char(36);not null;constraint:OnDelete:CASCADE"`
	Userid       string     `gorm:"type:char(36);not null;constraint:OnDelete:CASCADE"`
	Borroweddate time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Returndate   *time.Time `gorm:"default:null"`
}

type BookTransaction struct {
	Userid        string    `gorm:"type:char(36);default:null"`
	Bookid        string    `gorm:"type:char(36);default:null"`
}