package config

import (
	pb "client/pb/bookpb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func InitBookServiceClient() (*grpc.ClientConn, pb.BookServiceClient) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(os.Getenv("BOOK_SERVICE_URI"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn, pb.NewBookServiceClient(conn)
}