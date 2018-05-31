package routing

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "fmt"
)

// BaseRouter returns only business valuable routes
func BaseRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet, http.MethodPost)
    return r
}

func homeHandler() func(http.ResponseWriter, *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        log.Printf("Request is processing... %s", r.URL.Path)
        fmt.Fprint(w, "Hello! Your request was processed!")
    }
}