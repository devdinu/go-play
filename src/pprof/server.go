package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/ping", http.HandlerFunc(ping))
	http.HandleFunc("/leak", http.HandlerFunc(leak))
	http.HandleFunc("/slow_ping", http.HandlerFunc(slow))
	port := os.Getenv("PORT")
	fmt.Printf("starting server on %s", port)
	http.ListenAndServe(":"+port, nil)
}

func ping(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("handle")
	rw.Write([]byte("pong"))
}

func leak(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("leaking goroutine"))
	go func() {
		client := http.Client{
			Timeout: 0,
		}
		resp, err := client.Get("http://localhost:1234/slow_ping")
		if err != nil {
			log.Println(err)
			return
		}
		d, _ := ioutil.ReadAll(resp.Body)
		rw.Write(d)
	}()
}

// Imagine this as a separate server...
func slow(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 60)
	rw.Write([]byte("slow ... ping"))
}
