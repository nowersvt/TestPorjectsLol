package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: go run main.go <config-file>")
    }
    filename := os.Args[1]
    connStr,err := dbURL(filename)
    if err != nil {
		log.Fatalf("Error generating DB URL: %v", err)
	}
	fmt.Println("Connection String:", connStr)
    PrintDBTime()
    MainHandler()
}
