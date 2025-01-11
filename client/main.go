package main

import (
	"context"
	"net/http"

	"log"

	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	e := echo.New()
	token := "Bearer 1234"

	e.POST("/send-message", func(c echo.Context) error {
		var message map[string]string
		if err := c.Bind(&message); err != nil {
			return err
		}

		md := metadata.Pairs("authorization", token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Did'nt connect : %v", err)
		}

		defer conn.Close()

		// client := pb.NewAuthServiceClient(conn)

		// res, err := client.SayHello(ctx, &pb.HelloRequest{Name: message["message"]})
		// if err != nil {
		// 	log.Fatalf("error while create request %v", err)
		// }

		return c.JSON(http.StatusOK, map[string]string{
			"message": ctx.Err().Error(),
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}