package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealtChecker(t *testing.T) {
	h := http.HandlerFunc(Ping)

	server := httptest.NewServer(h)

	resp, _ := http.Get(server.URL + "/ping")
	d, _ := ioutil.ReadAll(resp.Body)
	response := string(d)
	assert.Equal(t, "pong", response)

	defer server.Close()
}
