package main

import (
	"net/http"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (cw *CustomResponseWriter) WriteHeader(status int) {
	cw.StatusCode = status
	cw.ResponseWriter.WriteHeader(status)
}
