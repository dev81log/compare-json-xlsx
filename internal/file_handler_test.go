package internal

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRenderHTML(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RenderHTML)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandleFileUpload(t *testing.T) {
	req, err := http.NewRequest("GET", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleFileUpload)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("jsonFile", "data.json")
	if err != nil {
		t.Fatal(err)
	}
	part.Write([]byte("test data"))

	part, err = writer.CreateFormFile("xlsxFile", "data.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	part.Write([]byte("test data"))

	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(HandleFileUpload)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status == http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	if _, err := os.Stat("../upload/data.json"); os.IsNotExist(err) {
		t.Error("file data.json not uploaded")
	}

	if _, err := os.Stat("../upload/data.xlsx"); os.IsNotExist(err) {
		t.Error("file data.xlsx not uploaded")
	}

}
