package repo

import "server/model"

type BookInterce interface {
	GetAllBooks(status string) ([]model.Book, error)
	ReturnBook(userId string, bookId string) error
	BorrowBook(userId string, bookId string) error
}
