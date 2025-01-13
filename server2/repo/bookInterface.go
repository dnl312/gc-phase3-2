package repo

import "server2/model"

type BookInterce interface {
	GetAllBooks(status string) ([]model.Book, error)
	ReturnBook(userId string, bookId string) error
	BorrowBook(userId string, bookId string) error
}
