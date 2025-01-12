package config

import (
	pb "client/pb/authpb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func InitUserServiceClient() (*grpc.ClientConn, pb.UserServiceClient) {
	// creds := credentials.NewTLS(&tls.Config{
	// 	InsecureSkipVerify: true,
	// })

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(os.Getenv("USER_SERVICE_URI"), opts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn, pb.NewUserServiceClient(conn)
}