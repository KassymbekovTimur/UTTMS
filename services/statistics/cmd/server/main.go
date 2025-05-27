package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"

	grpcAdapter "github.com/KassymbekovTimur/UTTMS/statistics/internal/adapter/grpc"
	natsAdapter "github.com/KassymbekovTimur/UTTMS/statistics/internal/adapter/nats"
	"github.com/KassymbekovTimur/UTTMS/statistics/internal/repository/postgres"
	"github.com/KassymbekovTimur/UTTMS/statistics/internal/usecase"
)

func main() {
	_ = godotenv.Load("../.env")

	dbURL := os.Getenv("DATABASE_URL")
	natsURL := os.Getenv("NATS_URL")
	grpcPort := os.Getenv("GRPC_PORT")
	// Connecting to DB
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL")

	repo := postgres.NewPostgresRepo(db)
	uc := usecase.NewStatsUsecase(repo)

	// NATS
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	natsAdapter.NewHandler(nc, uc)
	log.Println("NATS handler initialized")

	//gRPC
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpcAdapter.NewGRPCServer(uc)
	log.Println("GRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
