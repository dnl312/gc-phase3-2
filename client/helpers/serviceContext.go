package helpers

import (
	pb "client/pb/bookpb"
	"context"
	"os"
	"time"

	"google.golang.org/grpc"
)

//Creates a new context embedded with auth for gRPC services
// func NewServiceContext() (context.Context, context.CancelFunc, error) {
// 	token, err := SignJwtForGrpc()
// 	if err != nil {
// 		log.Printf("Error signing JWT for gRPC: %v", err)
// 		return nil, nil, echo.NewHTTPError(echo.ErrInternalServerError.Code, err.Error())
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
// 	return ctxWithAuth, cancel, nil
// }

func NewServiceContext() (context.Context, context.CancelFunc, pb.BookServiceClient, error) {
    conn, err := grpc.Dial(os.Getenv("BOOK_SERVICE_URI"), grpc.WithInsecure())
    if err != nil {
        return nil, nil, nil, err
    }
    client := pb.NewBookServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    return ctx, cancel, client, nil
}