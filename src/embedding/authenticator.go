package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func Authenticator(next http.HandlerFunc) http.HandlerFunc {
	authHandler := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		crw := CustomResponseWriter{ResponseWriter: w}
		if auth != "gopher" {
			crw.WriteHeader(http.StatusUnauthorized)
			crw.Write([]byte("{\"message\": \"Unauthorized\"}"))
			return
		}

		next(&crw, r)

		fmt.Printf("called: url: %s, statuscode: %d\n", r.URL.Path, crw.StatusCode)
	}

	bufio.Reader
	io.Reader
	return http.HandlerFunc(authHandler)
}
