package repo

import (
	"log"
	"server/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return BookRepository{
		DB: db,
	}
}

func (r BookRepository) GetAllBooks (status string) ([]model.Book, error) {
	var books []model.Book
	log.Printf("Status: %s", status)
	err := r.DB.Table("books_g2p3w2").Where("status = ?", status).Scan(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r BookRepository) BorrowBook (userId string, bookId string)  error {
	err := r.DB.Table("books_g2p3w2").Where("id = ?", bookId).Updates(map[string]interface{}{"user_id": userId, "status": "Borrowed"}).Error
	if err != nil {
		return err
	}

	var borrowedBook model.BorrowedBook
	borrowedBook.ID = uuid.New().String()
	borrowedBook.BookID = bookId
	borrowedBook.UserID = userId
	borrowedBook.BorrowedDate = time.Now()
	err = r.DB.Table("borrowedbooks_g2p3w2").Create(&borrowedBook).Error
	if err != nil {
		return err
	}

	return nil
}

func (r BookRepository) ReturnBook (userId string, bookId string) error {
	err := r.DB.Table("books_g2p3w2").Where("id = ?", bookId).Updates(map[string]interface{}{"user_id": nil, "status": "Available"}).Error
	if err != nil {
		return err
	}

	err = r.DB.Table("borrowedbooks_g2p3w2").Where("book_id = ? AND user_id = ? AND return_date IS NULL", bookId, userId).Updates(map[string]interface{}{"return_date": time.Now()}).Error
	if err != nil {
		return err
	}

	return nil
}