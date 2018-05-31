package routing_test

import (
    "testing"
    "github.com/gitAskur/gophercon/pkg/routing"
    "net/http/httptest"
    "net/http"
    "log"
)

func TestBaseRouter(t *testing.T) {
    handler := routing.BaseRouter()

    ts := httptest.NewServer(handler)
    defer ts.Close()
    res, err := http.Get(ts.URL + "/home")

    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode != http.StatusOK{
        t.Errorf("Wrong status code %d (%d expected)", res.StatusCode, http.StatusOK)
    }
}


