package internal

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/fileUpload.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/download", http.StatusFound)
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	r.ParseMultipartForm(32 << 20)

	jsonFile, handler, err := r.FormFile("jsonFile")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo json: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer jsonFile.Close()

	xlsxFile, handler2, err := r.FormFile("xlsxFile")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo xlsx: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer xlsxFile.Close()

	f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	j, err := os.OpenFile("./upload/"+handler2.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer j.Close()

	io.Copy(f, jsonFile)
	io.Copy(j, xlsxFile)

	ConverterFiles()
}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=relatorio.xlsx")
	http.ServeFile(w, r, "upload/relatorio.xlsx")
	fmt.Println("Download concluído")
}
