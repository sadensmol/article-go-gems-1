package main

import (
	"context"
	"log"
	"time"

	v1 "github.com/sadensmol/article-go-gems-1/api/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	cl := v1.NewAPIV1ServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_ = cancel
	println("calling Test ...")
	resp, err := cl.Test(ctx, &v1.TestRequest{WaitSec: 5, RequestDB: true})
	println(resp)
	if err != nil {
		log.Fatalf("Failed to call Test: %v", err)
	}

}
