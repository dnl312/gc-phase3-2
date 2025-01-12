package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"server/config"
	"server/model"
	"server/repo"

	pb "server/pb"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	pb.UnimplementedUserServiceServer
}

type BookServiceServer struct {
	pb.UnimplementedBookServiceServer
}

func (s *AuthServiceServer) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	
	
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	log.Printf("User object: %+v", user)

	token, err := repo.NewUserRepository(config.DB).LoginUser(user)
	if err != nil {
		debug := fmt.Sprintf("Login failed: %s", user.Username)
		return &pb.LoginResponse{Message: debug}, err
	}

	return &pb.LoginResponse{Message: token}, nil
}

func (s *AuthServiceServer) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := model.RegisterUser{
		Username: req.Username,
		Password: req.Password,
	}

	log.Printf("User object: %+v", user)

	err := repo.NewUserRepository(config.DB).RegisterUser(user)
	if err != nil {
		debug := fmt.Sprintf("Registration failed: %s", user.Username)
		return &pb.RegisterResponse{Message: debug}, err
	}

	return &pb.RegisterResponse{Message: "Registration successful"}, nil
}

func (b *BookServiceServer) GetAllBooks(ctx context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
    books, err := repo.NewBookRepository(config.DB).GetAllBooks(req.Status)
    if err != nil {
        log.Printf("Error fetching books (status=%s): %v", req.Status, err)
        return nil, status.Errorf(codes.Internal, "failed to fetch books: %v", err) // Return error specific to this method
    }

    var responseBooks []*pb.Book
    for _, book := range books {
        responseBooks = append(responseBooks, &pb.Book{
            Id:    book.ID,
            Title: book.Title,
            Status: book.Status,
        })
    }

    return &pb.GetAllBooksResponse{Books: responseBooks}, nil
}

func main() {
	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// e.POST("/users/login", service.LoginUser)
	// e.POST("/users/register", service.RegisterUser)

	config.InitDB()
	defer config.CloseDB()

	config.ClearPreparedStatements()

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &AuthServiceServer{})
	pb.RegisterBookServiceServer(s, &BookServiceServer{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}