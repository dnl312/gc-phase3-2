package controller

import (
	"context"
	"server/models"
	pb "server/pb"

	repository "server/repo"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	Repository repository.UserRepository
}

func NewServerController(repo repository.UserRepository) Server {
	return Server{Repository: repo}
}

func (s *Server) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := &models.User{
		Username: in.Username,
		Password: in.Password,
	}
	err := s.Repository.RegisterUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterUserResponse{Message: "User has been registered"}, nil
}

