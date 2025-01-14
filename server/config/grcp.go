package config

import (
	"log"
	"net"
	"os"
	"server/middleware"
	"server/pb"

	// grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

func ListenAndServeGrpc(controller pb.UserServiceServer) {
	port := os.Getenv("GRPC_PORT")
	
	lis, err := net.Listen("tcp", "0.0.0.0:" + port)
	if err != nil {
		log.Fatal(err)
	}
	
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middleware.NewInterceptorLogger()),
		),
	)
	pb.RegisterUserServiceServer(grpcServer, controller)
	//pb.RegisterUserServer(grpcServer, controller)

	log.Println("\033[36mGRPC server is running on port:", port, "\033[0m")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server gRPC:", err)
	}
}
