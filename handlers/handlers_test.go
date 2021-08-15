package handlers

import (
	"github.com/stretchr/testify/assert"
	"go-todos/apis"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodos(t *testing.T) {
	client := &apis.MockHttpClient{}
	client.On("Get", URL).Return([]byte(`[{"title":"delectus aut autem"}]`), nil)

	req, _ := http.NewRequest("GET", "/todos",nil)
	rec := httptest.NewRecorder()
	h := GetTodos(client)
	h.ServeHTTP(rec, req)
	assert.Contains(t, rec.Body.String(), `"title":"delectus aut autem"`)
	client.AssertExpectations(t)
}
