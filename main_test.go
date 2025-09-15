package main

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/stretchr/testify/assert"

  "demo/router"
)

func TestPingRoute(t *testing.T) {
  r := router.SetupRouter()

  w := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/hello", nil)
  r.ServeHTTP(w, req)

  assert.Equal(t, 200, w.Code)
  assert.Equal(t, "hello world", w.Body.String())
}
