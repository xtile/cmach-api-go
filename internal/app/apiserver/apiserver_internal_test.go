package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello! :)")
}

func TestAPIServer_HandleRoot(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	s.handleRoot().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Root is evil!")
}
