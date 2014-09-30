package login

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetLogin(t *testing.T) {
	request, _ := http.NewRequest("GET", "/login", nil)
	response := httptest.NewRecorder()
	GetLogin(response, request)
	t.Log(response)
}
