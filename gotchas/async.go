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
	http.HandleFunc("/panic", panicHandler(
		func(w http.ResponseWriter, r *http.Request) {
			panic("OMG Panic!")
		}))
	fmt.Println("serving requests...")
	http.ListenAndServe(":8080", nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	//This will result in context cancel error
	go getQuote(context.Background())
	fmt.Println("done...")
}

func getQuote(ctx context.Context) {
	url := "https://talaikis.com/api/quotes/random/"
	r, _ := http.NewRequest("GET", url, nil)
	// IMITATED PROCESSING TIME
	time.Sleep(2 * time.Second)

	r = r.WithContext(ctx)
	r.Close = true

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Oops! ", err)
		return
	}

	defer resp.Body.Close()

	quote := decode(resp.Body)
	fmt.Println("Quote of the day:  ", quote)
}

func decode(data io.Reader) string {
	var q quote
	if err := json.NewDecoder(data).Decode(&q); err != nil {
		return "Random-Quote"
	}
	return q.Quote
}

func panicHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				//REPORT THIS
				fmt.Println("recovered...", r)
			}
		}()
		//ACTUAL CALL
		next(w, r)
	}
}
