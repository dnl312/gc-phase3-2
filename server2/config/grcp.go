package config

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"server/middleware"
	"server/pb"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Println("No metadata found")
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	log.Printf("Metadata received: %v", md)

	authHeader, ok := md["authorization"]
    if !ok || len(authHeader) == 0 {
        return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
    }

	token := md["authorization"]

	parts := strings.Split(token[0], " ")
	tokenString := parts[1]

	
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, status.Errorf(http.StatusUnauthorized, "Invalid Authorization Header Format")
	}

	log.Println("Token validated successfully with value:", tokenString)
	return ctx, nil
}

func ListenAndServeGrpc(controller pb.BookServiceServer) {
	port := os.Getenv("GRPC_PORT")
	
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryAuthInterceptor),
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middleware.NewInterceptorLogger()),
		),
	)
	pb.RegisterBookServiceServer(grpcServer, controller)
	//pb.RegisterUserServer(grpcServer, controller)

	log.Println("\033[36mGRPC server is running on port:", port, "\033[0m")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server gRPC:", err)
	}
}
