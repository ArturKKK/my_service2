package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArturKKK/my_service2/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
