package repo

import (
	"server2/model"

	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
    mock.Mock
}

func NewMockBookRepository() MockBookRepository {
	return MockBookRepository{}
}

func (m *MockBookRepository) GetAllBooks(status string) ([]model.Book, error) {
    args := m.Called(status)
    return args.Get(0).([]model.Book), args.Error(1)
}

func (m *MockBookRepository) BorrowBook(userId string, bookId string) error {
    args := m.Called(userId, bookId)
    return args.Error(0)
}

func (m *MockBookRepository) ReturnBook(userId string, bookId string) error {
    args := m.Called(userId, bookId)
    return args.Error(0)
}

func (m *MockBookRepository) UpdateBookStatus() error {
    args := m.Called()
    return args.Error(0)
}