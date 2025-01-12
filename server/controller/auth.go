package controller

import (
	"context"
	"server/model"
	pb "server/pb"
	"server/repo"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	Repository repo.AuthInterface
}

func NewAuthController(r repo.AuthInterface) Server {
	return Server{
		Repository: r,
	}
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.Repository.LoginUser(model.User{Username: req.Username,Password:  req.Password})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Message: token}, nil
}

func (s *Server) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	
	err := s.Repository.RegisterUser(model.RegisterUser{Username: req.Username, Password:  req.Password})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}
