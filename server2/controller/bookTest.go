package controller

import (
	"context"
	"server2/model"
	pb "server2/pb"
	"server2/repo"
	"testing"
	"time"

	"github.com/stretchr/testify@v1.9.0/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	bookRepository = repo.NewMockBookRepository()
	bookController = NewBookController(&bookRepository)
)

var (
	modelBook = model.Book{
        ID: primitive.NewObjectID().Hex(),
		Title: "Book 1",
		Status: "Available",
        UserID: func(s string) *string { return &s }(primitive.NewObjectID().Hex()),
		PublishedDate: time.Now(),
		Author: "Author 1",
	}
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGetAllBooks(t *testing.T){
	var (
		status string = "Available"
	)

	books := []model.Book{
		modelBook,
	}

	bookRepository.Mock.On("GetAllBooks", context.Background(), status).Return(books, nil)

    pbResponse, err := bookController.GetAllBooks(context.Background(), &pb.GetAllBooksRequest{Status: status})
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

// func TestBorrowBook(t *testing.T) {
//     mockRepo := new(MockBookRepository)
//     server := NewBookController(mockRepo)

//     mockRepo.On("BorrowBook", "user1", "book1").Return(nil)

//     req := &pb.BorrowBookRequest{UserId: "user1", BookId: "book1"}
//     resp, err := server.BorrowBook(context.Background(), req)

//     assert.NoError(t, err)
//     assert.NotNil(t, resp)
//     assert.Equal(t, "Book borrowed successfully", resp.Message)
//     mockRepo.AssertExpectations(t)
// }

// func TestReturnBook(t *testing.T) {
//     mockRepo := new(MockBookRepository)
//     server := NewBookController(mockRepo)

//     mockRepo.On("ReturnBook", "user1", "book1").Return(nil)

//     req := &pb.ReturnBookRequest{UserId: "user1", BookId: "book1"}
//     resp, err := server.ReturnBook(context.Background(), req)

//     assert.NoError(t, err)
//     assert.NotNil(t, resp)
//     assert.Equal(t, "Book returned successfully", resp.Message)
//     mockRepo.AssertExpectations(t)
// }

// func TestUpdateBookStatus(t *testing.T) {
//     mockRepo := new(MockBookRepository)
//     server := NewBookController(mockRepo)

//     mockRepo.On("UpdateBookStatus", "book1", "Available").Return(nil)

//     req := &pb.UpdateBookStatusRequest{}
//     resp, err := server.UpdateBookStatus(context.Background(), req)

//     assert.NoError(t, err)
//     assert.NotNil(t, resp)
//     assert.Equal(t, "Task successfully", resp.Message)
//     mockRepo.AssertExpectations(t)
// }