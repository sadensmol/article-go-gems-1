package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"time"

	_ "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"

	v1 "github.com/sadensmol/article-go-gems-1/api/v1"
	"github.com/sadensmol/article-go-gems-1/internal/config"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.GetConfig()
	db, err := sql.Open("postgres", cfg.ConnectionURL())
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register your service implementation with the gRPC server
	v1.RegisterAPIV1ServiceServer(grpcServer, &Server{db: db})

	// Create a TCP listener on port 5050
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	// Start the gRPC server
	log.Println("Starting gRPC server on port 5050...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

	// setup db
}

type Server struct {
	v1.UnimplementedAPIV1ServiceServer
	db *sql.DB
}

func (s *Server) Test(ctx context.Context, in *v1.TestRequest) (*v1.TestResponse, error) {
	log.Println("received test request ...")

	var tx *sql.Tx
	var err error
	if in.RequestDB {
		log.Println("requesting DB ...")
		tx, err = s.db.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal("failed to begin transaction", err)
		}
		tx.ExecContext(ctx, "SELECT 1")
		defer func() {
			log.Println("rolling back transaction ...")
			err = tx.Rollback()
			if err != nil {
				log.Println("failed to rollback transaction", err)
			}
		}()
	}

	select {
	case <-ctx.Done():
		log.Println("context is done...")
	case <-time.After(time.Duration(in.WaitSec) * time.Second):
		log.Println("waited for", in.WaitSec, "seconds")
	}

	if in.RequestDB {
		log.Println("committing transaction ...")
		err := tx.Commit()
		if err != nil {
			log.Println("failed to commit transaction", err)
		}
	}
	return &v1.TestResponse{}, nil
}
