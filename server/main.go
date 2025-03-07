package main

import (
	"log"
	"server/config"
	"server/controller"
	"server/repo"

	"github.com/joho/godotenv"
)

// type AuthServiceServer struct {
// 	pb.UnimplementedUserServiceServer
// }

// type BookServiceServer struct {
// 	pb.UnimplementedBookServiceServer
// }

// // func (s *AuthServiceServer) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

// // 	user := model.User{
// // 		Username: req.Username,
// // 		Password: req.Password,
// // 	}

// // 	log.Printf("User object: %+v", user)

// // 	token, err := repo.NewUserRepository(config.DB).LoginUser(user)
// // 	if err != nil {
// // 		debug := fmt.Sprintf("Login failed: %s", user.Username)
// // 		return &pb.LoginResponse{Message: debug}, err
// // 	}

// // 	return &pb.LoginResponse{Message: token}, nil
// // }

// // func (s *AuthServiceServer) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
// // 	user := model.RegisterUser{
// // 		Username: req.Username,
// // 		Password: req.Password,
// // 	}

// // 	log.Printf("User object: %+v", user)

// // 	err := repo.NewUserRepository(config.DB).RegisterUser(user)
// // 	if err != nil {
// // 		debug := fmt.Sprintf("Registration failed: %s", user.Username)
// // 		return &pb.RegisterResponse{Message: debug}, err
// // 	}

// // 	return &pb.RegisterResponse{Message: "Registration successful"}, nil
// // }

// func (b *BookServiceServer) GetAllBooks(ctx context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
//     books, err := repo.NewBookRepository(config.DB).GetAllBooks(req.Status)
//     if err != nil {
//         log.Printf("Error fetching books (status=%s): %v", req.Status, err)
//         return nil, status.Errorf(codes.Internal, "failed to fetch books: %v", err) // Return error specific to this method
//     }

//     var responseBooks []*pb.Book
//     for _, book := range books {
//         responseBooks = append(responseBooks, &pb.Book{
//             Id:    book.ID,
//             Title: book.Title,
//             Status: book.Status,
//         })
//     }

//     return &pb.GetAllBooksResponse{Books: responseBooks}, nil
// }

// func (b *BookServiceServer) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
// 	err := repo.NewBookRepository(config.DB).BorrowBook(req.UserId, req.BookId)
// 	if err != nil {
// 		log.Printf("Error borrowing book (userId=%s, bookId=%s): %v", req.UserId, req.BookId, err)
// 		return nil, status.Errorf(codes.Internal, "failed to borrow book: %v", err) // Return error specific to this method
// 	}

// 	return &pb.BorrowBookResponse{Message: "Book borrowed successfully"}, nil
// }

// func (b *BookServiceServer) ReturnBook(ctx context.Context, req *pb.ReturnBookRequest) (*pb.ReturnBookResponse, error) {
// 	err := repo.NewBookRepository(config.DB).ReturnBook(req.UserId, req.BookId)
// 	if err != nil {
// 		log.Printf("Error returning book (userId=%s, bookId=%s): %v", req.UserId, req.BookId, err)
// 		return nil, status.Errorf(codes.Internal, "failed to return book: %v", err) // Return error specific to this method
// 	}

// 	return &pb.ReturnBookResponse{Message: "Book returned successfully"}, nil
// }

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

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	config.ClearPreparedStatements()
	
	defer config.CloseDB()
	userRepository := repo.NewUserRepository(db)
	userController := controller.NewAuthController(userRepository)

	config.ListenAndServeGrpc(&userController)
	// config.ClearPreparedStatements()

	// grpcPort := os.Getenv("GRPC_PORT")
	// if grpcPort == "" {
	// 	grpcPort = "50051"
	// }

	// lis, err := net.Listen("tcp", ":"+grpcPort)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// s := grpc.NewServer()
	// pb.RegisterUserServiceServer(s, &AuthServiceServer{})
	// pb.RegisterBookServiceServer(s, &BookServiceServer{})

	// log.Printf("gRPC server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}