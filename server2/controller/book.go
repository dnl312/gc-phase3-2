package controller

import (
	"context"
	"log"
	"server2/config"
	pb "server2/pb"
	"server2/repo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedBookServiceServer
	Repository repo.BookRepository
}

func NewBookController(r repo.BookInterce) Server {
	return Server{
		Repository: r.(repo.BookRepository),
	}
}

func (b *Server) GetAllBooks(ctx context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	log.Printf("Fetching books with status: %s", req.Status)
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

func (b *Server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	err := repo.NewBookRepository(config.DB).BorrowBook(req.UserId, req.BookId)
	if err != nil {
		log.Printf("Error borrowing book (userId=%s, bookId=%s): %v", req.UserId, req.BookId, err)
		return nil, status.Errorf(codes.Internal, "failed to borrow book: %v", err) // Return error specific to this method
	}

	return &pb.BorrowBookResponse{Message: "Book borrowed successfully"}, nil
}

func (b *Server) ReturnBook(ctx context.Context, req *pb.ReturnBookRequest) (*pb.ReturnBookResponse, error) {
	err := repo.NewBookRepository(config.DB).ReturnBook(req.UserId, req.BookId)
	if err != nil {
		log.Printf("Error returning book (userId=%s, bookId=%s): %v", req.UserId, req.BookId, err)
		return nil, status.Errorf(codes.Internal, "failed to return book: %v", err) // Return error specific to this method
	}

	return &pb.ReturnBookResponse{Message: "Book returned successfully"}, nil
}
