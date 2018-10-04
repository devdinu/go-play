package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("running....")
	http.ListenAndServe(":8080", router())
}

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)
	mux.Handle("/coolping", HealthChecker(service{}))
	return mux
}

type service struct{}

func (s service) Ping() error { return nil }
