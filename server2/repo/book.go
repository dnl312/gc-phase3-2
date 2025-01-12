package repo

import (
	"fmt"
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
	    // Check if the user exists
    var user model.User
    err := r.DB.Table("users_g2p3w2").Where("id = ?", userId).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return fmt.Errorf("user does not exist")
        }
        return err
    }

	var book model.Book
	err = r.DB.Table("books_g2p3w2").Where("id = ?", bookId).First(&book).Error
	if err != nil {
		return err
	}

	if book.Status != "Available" {
		return fmt.Errorf("book is not available for borrowing")
	}


	err = r.DB.Table("books_g2p3w2").Where("id = ?", bookId).Updates(map[string]interface{}{"userid": userId, "status": "Borrowed"}).Error
	if err != nil {
		return err
	}

	var borrowedBook model.BorrowedBook
	borrowedBook.ID = uuid.New().String()
	borrowedBook.Bookid = bookId
	borrowedBook.Userid = userId
	borrowedBook.Borroweddate = time.Now()

	err = r.DB.Table("borrowedbooks_g2p3w2").Create(&borrowedBook).Error
	if err != nil {
		return err
	}

	return nil
}

func (r BookRepository) ReturnBook (userId string, bookId string) error {
	log.Printf("ReturnBook1: %s %s", userId, bookId)
	var user model.User
    err := r.DB.Table("users_g2p3w2").Where("id = ?", userId).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return fmt.Errorf("user does not exist")
        }
        return err
    }
	
	err = r.DB.Table("books_g2p3w2").Where("id = ?", bookId).Updates(map[string]interface{}{"userid": nil, "status": "Available"}).Error
	if err != nil {
		return err
	}

	log.Printf("ReturnBook2: %s %s", userId, bookId)
	err = r.DB.Table("borrowedbooks_g2p3w2").Where("bookid = ? AND userid = ? AND returndate IS NULL", bookId, userId).Updates(map[string]interface{}{"returndate": time.Now()}).Error
	if err != nil {
		return err
	}

	return nil
}