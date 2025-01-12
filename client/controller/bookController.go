package controller

import (
	"client/helpers"
	"client/model"
	pb "client/pb/bookpb"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	Client pb.BookServiceClient
}

func NewBookController(client pb.BookServiceClient) BookController {
	return BookController{
		Client: client,
	}
}

func (b BookController) GetAllBooks(ctx echo.Context) error{
		status := ctx.QueryParam("status")
		if status == "" {
			status = "Available"
		}

		serviceCtx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		r, err := b.Client.GetAllBooks(serviceCtx, &pb.GetAllBooksRequest{Status: status})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "could not get books"})
		}
		log.Printf("Books Response: %v", r.GetBooks())
		
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"books": r.Books,
		})
}

func (b BookController) BorrowBook(ctx echo.Context) error{
		var req model.BookTransaction
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}
		
		serviceCtx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		r, err := b.Client.BorrowBook(serviceCtx, &pb.BorrowBookRequest{UserId: req.UserID, BookId: req.BookID})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "could not borrow book"})
		}
		log.Printf("Borrow Response: %s", r.GetMessage())

		return ctx.JSON(http.StatusOK, map[string]string{
			"message": r.Message,
		})
}

func (b BookController) ReturnBook(ctx echo.Context) error{
		var req model.BookTransaction
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}
		log.Printf("ReturnBook1: %s %s", req.UserID, req.BookID)
		serviceCtx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		r, err := b.Client.ReturnBook(serviceCtx, &pb.ReturnBookRequest{UserId: req.UserID, BookId: req.BookID})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "could not return book"})
		}	
		log.Printf("Return Response: %s", r.GetMessage())
	
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": r.Message,
		})
}
