package main

import "fmt"

func dbURL(filename string) (string, error) {
    config, err := LoadConfig(filename)
    if err != nil {
        return "", err
    }

    fmt.Printf("Loaded Config: %+v\n", config)

    host := "localhost"
    user := config.Postgres.Environment.User
    password := config.Postgres.Environment.Password
    dbName := config.Postgres.Environment.DbName
    port, err := extractPort(config.Postgres.Ports)
    if err != nil {
        return "", err
    }

    fmt.Println("Host:", host)
    fmt.Println("User:", user)
    fmt.Println("Password:", password)
    fmt.Println("Database:", dbName)
    fmt.Println("Port:", port)

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName)
    return connStr, nil
}