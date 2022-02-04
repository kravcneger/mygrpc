package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/kravcneger/mygrpc/internal"
	pb "github.com/kravcneger/mygrpc/mygrpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
	ctx  = context.Background()
)

type server struct {
	pb.UnimplementedMyGrpcServer
	db    internal.Database
	redis *redis.Client
	cache *cache.Cache
}

func (s *server) CreateUser(ctx context.Context, user *pb.User) (*pb.StatusCode, error) {
	err := s.db.CreateUser(user.Login, user.Email)
	if err != nil {
		return &pb.StatusCode{}, err
	}
	return &pb.StatusCode{Code: 200}, nil
}

func (s *server) ListUsers(rect *pb.Query, stream pb.MyGrpc_ListUsersServer) error {
	users := make([]internal.User, 0)
	data, err := s.redis.Get(ctx, "list").Result()

	if err != nil && err != redis.Nil {
		return err
	}
	// if chache does not have any data get it from database
	if err == redis.Nil {
		users, err = s.db.GetUsers()
		if err != nil {
			return err
		}
		// add data to cache
		json, err := json.Marshal(&users)
		if err != nil {
			return err
		}
		if _, err := s.redis.Set(ctx, "list", json, time.Minute).Result(); err != nil {
			return err
		}
	} else {
		err = json.Unmarshal([]byte(data), &users)
		if err != nil {
			return err
		}
	}

	for _, user := range users {
		var protoUser *pb.User
		protoUser = userToProtoUser(user)
		if err := stream.Send(protoUser); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) DeleteUser(ctx context.Context, user *pb.User) (*pb.StatusCode, error) {
	err := s.db.DeleteUser(int(user.Id))
	if err != nil {
		return &pb.StatusCode{}, err
	}
	return &pb.StatusCode{Code: 200}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	serv := server{}

	// Initialize redis connection
	serv.redis = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	defer serv.redis.Close()

	// Initialize chache
	serv.cache = cache.New(&cache.Options{
		Redis:      serv.redis,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	// Initialize postgress connection
	dbUser, dbPassword, dbPort, dbName := os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB")
	serv.db, err = internal.Initialize(dbUser, dbPassword, dbPort, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v, dbUser - %s,	dbPassword - %s, dbPort - %s, dbName - %s",
			err, dbUser, dbPassword, dbPort, dbName)
	}
	defer serv.db.Conn.Close()

	pb.RegisterMyGrpcServer(s, &serv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func userToProtoUser(user internal.User) *pb.User {
	return &pb.User{
		Id:    int64(user.Id),
		Login: user.Login,
		Email: user.Email,
	}
}
