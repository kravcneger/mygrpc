package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/kravcneger/mygrpc/mygrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

type Client struct {
	pb.MyGrpcClient
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func createUser(client pb.MyGrpcClient) {
	flag.Parse()
	if len(os.Args) < 4 {
		log.Fatalf("fatal: add login and email params")
	}
	login := os.Args[2]
	email := os.Args[3]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.CreateUser(ctx, &pb.User{Login: login, Email: email})
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}
	log.Printf("User is created: %d", status.Code)
}

func deleteUser(client pb.MyGrpcClient) {
	flag.Parse()
	if len(os.Args) < 3 {
		log.Fatalf("fatal: add id param")
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("fatal: incorrect Id param")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.DeleteUser(ctx, &pb.User{Id: int64(id)})
	if err != nil {
		log.Fatalf("could not delete: %v", err)
	}
	log.Printf("User is deleted: %d", status.Code)
}

func listUsers(client pb.MyGrpcClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	stream, err := client.ListUsers(ctx, &pb.Query{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User id: %d, login: %s, email: %s", user.Id, user.Login, user.Email)
		fmt.Println()
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyGrpcClient(conn)

	if len(os.Args) < 2 {
		log.Fatalf("fatal: you should specify the command")
	}

	command := os.Args[1]
	switch command {
	case "create":
		createUser(c)
	case "list":
		listUsers(c)
	case "delete":
		deleteUser(c)
	default:
		log.Fatalf("unknown comand: %v", command)
	}

}

func Initialize(username, password, port, database string) (*sql.DB, error) {
	var db *sql.DB
	dsn := fmt.Sprint("user=postgres password=postgres dbname=mygrpc host=localhost port=5433 sslmode=disable")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db = conn
	err = db.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
