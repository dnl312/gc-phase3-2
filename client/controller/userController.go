package controller

import (
	"client/model"
	pb "client/pb/authpb"
	"context"
	"time"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Client pb.UserServiceClient
}

func NewUserController(client pb.UserServiceClient) UserController {
	return UserController{
		Client: client,
	}
}

func (u UserController) LoginUser (ctx echo.Context) error{
		var req model.LoginRequest
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}

		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
		serviceCtx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		 log.Printf("could not login: %v", req)

		r, err := u.Client.LoginUser(serviceCtx, &pb.LoginRequest{Username: req.Username, Password: req.Password})
		if err != nil {
			log.Printf("could not login2: %v", err)
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "login error"})
		}
		log.Printf("Login Response: %s", r.GetMessage())

		return ctx.JSON(http.StatusOK, map[string]string{
			"message": r.Message,
		})
}

func (u UserController) RegisterUser (ctx echo.Context) error{
	var req model.RegisterUser
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	serviceCtx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := u.Client.RegisterUser(serviceCtx, &pb.RegisterRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		log.Printf("could not register: %v", err)
		log.Printf("could not register: %v", req)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "registration failed"})
	}
	log.Printf("Register Response: %s", r.GetMessage())

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": r.Message,
	})
}