package main

import (
    "net/http"
    "log"

    "github.com/gitAskur/gophercon/pkg/routing"
)
// go run ./cmd/gophercon/gophercon.go
func main() {
    log.Printf("Service is starting...")

    r := routing.BaseRouter()

    log.Fatal(http.ListenAndServe(":8000", r))

}

