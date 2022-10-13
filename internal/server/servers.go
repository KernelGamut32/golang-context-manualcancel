package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func SlowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		// time.Sleep(2 * time.Second)
		// time.Sleep(6 * time.Second)
		// w.Write([]byte("Slow response"))

		ctx := r.Context()
		select {
		case <-ctx.Done():
			fmt.Println("server shut down")
			return
		case <-time.After(6 * time.Second):
			w.Write([]byte("Slow response"))
		}
	}))
	return s
}

func FastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}))
	return s
}
