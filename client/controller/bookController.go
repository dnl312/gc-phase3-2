package controller

import (
	"client/model"
	pb "client/pb/bookpb"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

type BookController struct {
	Client pb.BookServiceClient
}

func NewBookController(client pb.BookServiceClient) BookController {
	return BookController{
		Client: client,
	}
}

// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept  json
// @Produce  json
// @Param status query string false "Book status"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (b BookController) GetAllBooks(ctx echo.Context) error{

	   	// claims, err := helpers.GetClaims(ctx)
		// if err != nil {
		// 	log.Printf("Error getting claims: %v", err)
		// 	return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization cookie does not exist"})
		// }

		// log.Printf("Authorization claims: %v", claims)
		// // Check for the authorization cookie
		// cookie, err := ctx.Cookie("Authorization")
		// if err != nil {
		// 	log.Printf("Authorization cookie does not exist: %v", err)
		// 	return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization cookie does not exist"})
		// }

		// print cookie.Value to log
		// log.Printf("Authorization cookie value: %v", claimed)

		// // Use the authorization cookie to make the gRPC request
		// req := &pb.GetAllBooksRequest{
		// 	Authorization: cookie.Value,
		// }

		status := ctx.QueryParam("status")
		if status == "" {
			status = "Available"
		}

		token := ctx.Request().Header.Get("Authorization")
		md := metadata.Pairs("Authorization", token)
		newCtx := metadata.NewOutgoingContext(context.Background(), md)

		serviceCtx, cancel := context.WithTimeout(newCtx, 10*time.Second)
		defer cancel()

		r, err := b.Client.GetAllBooks(serviceCtx, &pb.GetAllBooksRequest{Status: status})
		if err != nil {
			log.Printf("Error fetching books (status=%s): %v", status, err)
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "could not get books"})
		}
		log.Printf("Books Response: %v", r.GetBooks())
		
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"books": r.Books,
		})
}

// @Summary Borrow a book
// @Description Borrow a book
// @Tags books
// @Accept  json
// @Produce  json
// @Param request body model.BookTransaction true "Book transaction details"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string
//@failure 500 {object} map[string]string
// @Router /books/borrow [post]
func (b BookController) BorrowBook(ctx echo.Context) error{
		var req model.BookTransaction
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}

		token := ctx.Request().Header.Get("Authorization")
		md := metadata.Pairs("Authorization", token)
		newCtx := metadata.NewOutgoingContext(context.Background(), md)

		serviceCtx, cancel := context.WithTimeout(newCtx, 10*time.Second)
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

//@summary Return a book
//@description Return a book
//@tags books
//@accept json
//@produce json
//@param request body model.BookTransaction true "User ID and Book ID"
//@security ApiKeyAuth
//@success 200 {object} map[string]string
//@failure 500 {object} map[string]string
//@router /books/return [post]
func (b BookController) ReturnBook(ctx echo.Context) error{
		var req model.BookTransaction
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}
		log.Printf("ReturnBook1: %s %s", req.UserID, req.BookID)

		token := ctx.Request().Header.Get("Authorization")
		md := metadata.Pairs("Authorization", token)
		newCtx := metadata.NewOutgoingContext(context.Background(), md)

		serviceCtx, cancel := context.WithTimeout(newCtx, 10*time.Second)
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

//@summary Update book status
//@description Update book status
//@tags books
//@accept json
//@produce json
//@success 200 {object} map[string]string
//@router /books/update-status [get]
func (b BookController) UpdateBookStatus(ctx echo.Context) error{

	serviceCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

	r, err := b.Client.UpdateBookStatus(serviceCtx, &pb.UpdateBookStatusRequest{})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "could not update book status"})
		}	
		log.Printf("Update Status Response: %s", r.GetMessage())
	
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": r.Message,
		})
	}