package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHome(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/home", nil)
	s.handleHome().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
