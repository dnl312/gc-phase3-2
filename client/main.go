package main

import (
	"log"

	"client/config"
	"client/controller"
	"client/helpers"
	"client/router"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Library API
// @version 1.0
// @description This is a simple library API
// @host localhost:8080
// @BasePath /
func main() {
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// if err != nil {
	// 	log.Printf("Didn't connect: %v", err)
	// }
	// defer conn.Close()
	// client := pb.NewUserServiceClient(conn)
	// clientBook := pbBook.NewBookServiceClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

////////////////////////////////////////
	e := echo.New()
	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}
	e.Use(middleware.Logger(), middleware.Recover(), middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	authClientConn, authClient := config.InitUserServiceClient()
	defer authClientConn.Close()

	bookClientConn, bookClient := config.InitBookServiceClient()
	defer bookClientConn.Close()

	userController := controller.NewUserController(authClient)
	bookController := controller.NewBookController(bookClient)

	router.Echo(e, userController, bookController)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

////////////////////////////////////////
	
	

	// e.POST("/users/login", func(c echo.Context) error {
	// 	var req model.LoginRequest
	// 	if err := c.Bind(&req); err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	// 	}
	// 	r, err := client.LoginUser(ctx, &pb.LoginRequest{Username: req.Username, Password: req.Password})
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "login error"})
	// 	}
	// 	log.Printf("Login Response: %s", r.GetMessage())

	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": r.Message,
	// 	})
	// })
	// e.POST("/users/register", func(c echo.Context) error {
	// 	var req model.RegisterUser
	// 	if err := c.Bind(&req); err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	// 	}
	// 	r, err := client.RegisterUser(ctx, &pb.RegisterRequest{Username: req.Username, Password: req.Password})
	// 	if err != nil {
	// 		log.Printf("could not register: %v", err)
	// 		log.Printf("could not register: %v", req)
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "registration failed"})
	// 	}
	// 	log.Printf("Register Response: %s", r.GetMessage())

	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": r.Message,
	// 	})
	// })

	// a := e.Group("/books")
	// a.GET("", func(c echo.Context) error {
	// 	status := c.QueryParam("status")
	// 	if status == "" {
	// 		status = "Available"
	// 	}
	// 	r, err := clientBook.GetAllBooks(ctx, &pbBook.GetAllBooksRequest{Status: status})
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not get books"})
	// 	}
	// 	log.Printf("Books Response: %v", r.GetBooks())
		
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"books": r.Books,
	// 	})
	// })
	// a.POST("/borrow", func(c echo.Context) error {
	// 	var req model.BookTransaction
	// 	if err := c.Bind(&req); err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	// 	}
	// 	r, err := clientBook.BorrowBook(ctx, &pbBook.BorrowBookRequest{UserId: req.UserID, BookId: req.BookID})
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not borrow book"})
	// 	}
	// 	log.Printf("Borrow Response: %s", r.GetMessage())

	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": r.Message,
	// 	})
	// })
	// a.POST("/return", func(c echo.Context) error {
	// 	var req model.BookTransaction
	// 	if err := c.Bind(&req); err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	// 	}
	// 	log.Printf("ReturnBook1: %s %s", req.UserID, req.BookID)
	// 	r, err := clientBook.ReturnBook(ctx, &pbBook.ReturnBookRequest{UserId: req.UserID, BookId: req.BookID})
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not return book"})
	// 	}	
	// 	log.Printf("Return Response: %s", r.GetMessage())
	
	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": r.Message,
	// 	})
	// })

	e.Logger.Fatal(e.Start(":8080"))
}