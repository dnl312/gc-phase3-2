package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dnl312/gc-phase3-2/server/pb"

	"google.golang.org/grpc"
)

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewUserServiceClient(conn)

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  r, err := c.LoginUser(ctx, &pb.LoginRequest{Username: "test", Password: "password"})
  if err != nil {
    log.Fatalf("could not login: %v", err)
  }
  log.Printf("Login Response: %s", r.GetMessage())

  r2, err := c.RegisterUser(ctx, &pb.RegisterRequest{Username: "test", Password: "password"})
  if err != nil {
    log.Fatalf("could not register: %v", err)
  }
  log.Printf("Register Response: %s", r2.GetMessage())
}