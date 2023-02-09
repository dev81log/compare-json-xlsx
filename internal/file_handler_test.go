package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHtmlHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HtmlHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código de status errado: got %v want %v",
			status, http.StatusOK)
	}
	if err != nil {
		t.Errorf("erro retornado do ParseFiles: %v", err)
	}

}
