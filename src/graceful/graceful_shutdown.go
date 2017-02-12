package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Ponged...")
	w.Write([]byte("pong"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", PingHandler)
	server := &http.Server{Addr: ":8888", Handler: mux}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go stopServerOnInterrupt(quit, server)
	log.Println("serving requests...")
	server.ListenAndServe()
}

// START OMIT
func stopServerOnInterrupt(quit chan os.Signal, server *http.Server) {
	value := <-quit
	log.Println("Signal Received:", value)
	server.Shutdown(context.Background()) // HL
	os.Exit(1)
}

//END OMIT
