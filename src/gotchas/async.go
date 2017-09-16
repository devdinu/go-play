package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func asyncServe() {
	http.HandleFunc("/ping", pingHandler)
	fmt.Println("serving requests...")
	http.ListenAndServe(":8080", nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	getQuote(r.Context())
	fmt.Println("done...")
}

func getQuote(ctx context.Context) {
	url := "https://talaikis.com/api/quotes/random/"
	r, _ := http.NewRequest("GET", url, nil)
	time.Sleep(3 * time.Second)
	r = r.WithContext(ctx)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Oops! ", err)
		return
	}
	quote := decode(resp.Body)
	fmt.Println("Quote of the day:  ", quote)
}

type quote struct {
	Quote  string
	Author string
	Cat    string
}

func decode(data io.Reader) string {
	var q quote
	if err := json.NewDecoder(data).Decode(&q); err != nil {
		return "Random-Quote"
	}
	return q.Quote
}
