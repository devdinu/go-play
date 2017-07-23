package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (cw *CustomResponseWriter) WriteHeader(status int) {
	cw.StatusCode = status
	cw.ResponseWriter.WriteHeader(status)
	io.ReadCloser
	ioutil.NopCloser
}
