package internal

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./templates/fileUpload.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmp.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	r.ParseMultipartForm(32 << 20)

	jsonFile, jsonHandler, err := r.FormFile("jsonFile")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo json: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer jsonFile.Close()

	xlsxFile, xlsxHandler, err := r.FormFile("xlsxFile")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo xlsx: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer xlsxFile.Close()

	jsonOpen, err := os.OpenFile("./upload/"+jsonHandler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonOpen.Close()

	xlsxOpen, err := os.OpenFile("./upload/"+xlsxHandler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer xlsxOpen.Close()

	io.Copy(jsonOpen, jsonFile)
	io.Copy(xlsxOpen, xlsxFile)

	defer ConverterFiles()                             // call converter function
	http.Redirect(w, r, "/download", http.StatusFound) // redirect to download page
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=relatorio.xlsx")
	http.ServeFile(w, r, "upload/relatorio.xlsx")

	fmt.Println("Download concluído")
}
