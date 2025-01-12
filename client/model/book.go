package model

type BookTransaction struct {
	UserID        string    `gorm:"type:char(36);default:null"`
	BookID        string    `gorm:"type:char(36);default:null"`
}
