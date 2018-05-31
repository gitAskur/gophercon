package cmd

import (
    "net/http"
    "fmt"
    "log"

    "github.com/gorilla/mux"
)

func main() {
    log.Printf("Service is starting...")

    http.HandleFunc("/home", homeHandler())
    r := mux.NewRouter()
    r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet, http.MethodPost)
}

func homeHandler() func(http.ResponseWriter, *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        log.Printf("Request is processing... %s", r.URL.Path)
        fmt.Fprint(w, "Hello! Your request was processed!")
    }
}

/*
func homeHandler() func(http.ResponseWriter, *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        log.Printf("Request is processing... %s", r.URL.Path)
        fmt.Fprint(w, "Hello! Your request was processed!")
    }
}
*/