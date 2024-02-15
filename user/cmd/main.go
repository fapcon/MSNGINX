package main

import (
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	"user/internal/grpc/user"
	"user/internal/repository"
	"user/internal/service"
)

func main() {
	time.Sleep(2 * time.Second)
	dbHost := "db"
	dbPort := "5432"
	dbUser := "userpostgres"
	dbPassword := "password"
	dbName := "userserv"
	sslmode := "disable"

	connectionString := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + sslmode

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	if err != nil {
		log.Fatalf("ping:%v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, email VARCHAR(255) UNIQUE NOT NULL, hashedpassword VARCHAR(255) NOT NULL)`)

	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo)

	serviceUser := user.NewServiceUser(userService)

	lis, err := net.Listen("tcp", ":44971")
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", "50053", err)
	}
	grpcServer := grpc.NewServer()

	userpr.RegisterUserServiceServer(grpcServer, serviceUser)

	log.Print("Starting gRPC server user...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
