package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "postgres"
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func InitializePostgres(username, password, port, database string) (*Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username, password, HOST, port, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return &db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return &db, err
	}
	log.Println("Database connection established")
	return &db, nil
}
