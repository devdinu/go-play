package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPing(t *testing.T) {

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/blaping", nil)

	Ping(w, r)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestDbError(t *testing.T) {
	c := new(checker)
	c.On("Ping").Times(1).Return(errors.New("someerr"))

	h := HealthChecker(c)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/bla", nil)

	h.ServeHTTP(w, r)

	c.AssertExpectations(t)
	assert.Equal(t, 503, w.Code)
	assert.Equal(t, "Service Unavailable", w.Body.String())
}

func TestDbSuccess(t *testing.T) {
	c := new(checker)
	c.On("Ping").Return(nil)

	h := HealthChecker(c)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/bla", nil)

	h.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

type checker struct{ mock.Mock }

func (c *checker) Ping() error {
	return c.Called().Error(0)
}
