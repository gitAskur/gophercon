package main

import (
    "net/http"
    "log"

    "github.com/gitAskur/gophercon/pkg/routing"
    "os"
)
// go run ./cmd/gophercon/gophercon.go
func main() {
    log.Printf("Service is starting...")

    port := os.Getenv("SERVICE_PORT")
    if len(port) == 0 {
        log.Fatal("Service Port not set!")
    }

    r := routing.BaseRouter()

    log.Fatal(http.ListenAndServe(":"+port, r))

}

