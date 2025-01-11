package config

import (
	"log"
	"net"
	"os"
	"server/controller"
	pb "server/pb"

	"google.golang.org/grpc"
)

func ListenAndServeGrpc(controller controller.Server) {
	port := os.Getenv("PORT")
	
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &controller)

	log.Println("\033[36mGRPC server is running on port:", port, "\033[0m")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server gRPC:", err)
	}
}