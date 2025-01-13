package controller

import (
	"context"
	"server2/model"
	pb "server2/pb"
	"server2/repo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	pbRequest := &pb.GetAllBooksRequest{
		Status: "Available",
	}

	books := []model.Book{
		modelBook,
	}

	bookRepository.Mock.On("GetAllBooks", pbRequest.Status).Return(books, nil)

	pbResponse, err := bookController.GetAllBooks(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}


func TestBorrowBook(t *testing.T) {
	pbRequest := &pb.BorrowBookRequest{
		UserId: primitive.NewObjectID().Hex(),
		BookId: primitive.NewObjectID().Hex(),
	}

	bookRepository.Mock.On("BorrowBook", pbRequest.UserId, pbRequest.BookId).Return(nil)


	pbResponse, err := bookController.BorrowBook(context.Background(), pbRequest)
	
    assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, "Book borrowed successfully", pbResponse.Message)
}

func TestReturnBook(t *testing.T) {
	pbRequest := &pb.ReturnBookRequest{
		UserId: primitive.NewObjectID().Hex(),
		BookId: primitive.NewObjectID().Hex(),
	}

	bookRepository.Mock.On("ReturnBook", pbRequest.UserId, pbRequest.BookId).Return(nil)

	pbResponse, err := bookController.ReturnBook(context.Background(), pbRequest)
	
    assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, "Book returned successfully", pbResponse.Message)
}

func TestUpdateBookStatus(t *testing.T) {
	pbRequest := &pb.UpdateBookStatusRequest{
	}
	
	bookRepository.Mock.On("UpdateBookStatus").Return(nil) 
    pbResponse, err := bookController.UpdateBookStatus(context.Background(), pbRequest)

    assert.NoError(t, err)
    assert.NotNil(t, pbResponse)
    assert.Equal(t, "Task successfully", pbResponse.Message)
}