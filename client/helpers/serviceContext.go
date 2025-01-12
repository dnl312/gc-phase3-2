package helpers

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	grpcMetadata "google.golang.org/grpc/metadata"
)

// Creates a new context embedded with auth for gRPC services
func NewServiceContext() (context.Context, context.CancelFunc, error) {
	token, err := SignJwtForGrpc()
	if err != nil {
		return nil, nil, echo.NewHTTPError(echo.ErrInternalServerError.Code, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	return ctxWithAuth, cancel, nil
}