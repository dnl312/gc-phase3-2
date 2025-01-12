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
