package main

import (
	"net/http"
)

type Checker interface {
	Ping() error
}

type HealthCheck struct {
	db Checker
}

func (hc *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := hc.db.Ping(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
		return
	}
	w.Write([]byte("pong"))
}

func HealthChecker(db Checker) http.Handler {
	return &HealthCheck{db: db}
}
