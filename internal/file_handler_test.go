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

func TestUploadHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler retornou código de status errado: got %v want %v",
			status, http.StatusFound)
	}
	if err != nil {
		t.Errorf("erro retornado do ParseFiles: %v", err)
	}

}

func TestDownloadHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/download", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DownloadHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código de status errado: got %v want %v",
			status, http.StatusOK)
	}
	if err != nil {
		t.Errorf("erro retornado do ParseFiles: %v", err)
	}

}
