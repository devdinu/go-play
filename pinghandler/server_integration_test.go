package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealtChecker(t *testing.T) {
	h := http.HandlerFunc(Ping)

	server := httptest.NewServer(h)

	resp, err := http.Get(server.URL + "/ping")

	require.NoError(t, err, "request failure")

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err, "reading response failure")
	assert.Equal(t, "pong", string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer server.Close()
}
