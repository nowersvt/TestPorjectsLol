package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	connStr := "postgres://test:12345@localhost:1337/testDB?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Connection Error: ", err)
	}
	defer conn.Close(context.Background())

	var now time.Time
	err = conn.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
	if err != nil{
		log.Fatal("Query execution error:", err)
	}

	fmt.Println("Current time in PostgreSQL", now)
}