package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/jackc/pgx/v5"
)

func GetDBTime(filename string) (time.Time, error) {
    connStr, err := dbURL(filename)
    if err != nil {
        return time.Time{}, err
    }

    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
        return time.Time{}, err
    }
    defer conn.Close(context.Background())

    var now time.Time
    err = conn.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
    if err != nil {
        return time.Time{}, err
    }
    return now, nil
}

func PrintDBTime() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: go run main.go <config-file>")
    }
    filename := os.Args[1]

    now, err := GetDBTime(filename)
    if err != nil {
        log.Fatal("Ошибка при получении времени из БД:", err)
    }
    fmt.Println("Текущее время в PostgreSQL:", now)
}