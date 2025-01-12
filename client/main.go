package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"client/model"
	"client/pb"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Didn't connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err = godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	e.POST("/users/login", func(c echo.Context) error {
		var req model.LoginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}
		r, err := client.LoginUser(ctx, &pb.LoginRequest{Username: req.Username, Password: req.Password})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "login error"})
		}
		log.Printf("Login Response: %s", r.GetMessage())

		return c.JSON(http.StatusOK, map[string]string{
			"message": r.Message,
		})
	})
	e.POST("/users/register", func(c echo.Context) error {
		var req model.RegisterUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}
		r, err := client.RegisterUser(ctx, &pb.RegisterRequest{Username: req.Username, Password: req.Password})
		if err != nil {
			log.Printf("could not register: %v", err)
			log.Printf("could not register: %v", req)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "registration failed"})
		}
		log.Printf("Register Response: %s", r.GetMessage())

		return c.JSON(http.StatusOK, map[string]string{
			"message": r.Message,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}