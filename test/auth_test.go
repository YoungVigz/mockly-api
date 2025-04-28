package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YoungVigz/mockly-api/internal/routes"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	router := routes.SetupTestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
