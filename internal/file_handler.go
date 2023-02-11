package internal

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./templates/fileUpload.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmp.Execute(w, nil)
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
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

	files := map[string]io.Reader{
		jsonHandler.Filename: jsonFile,
		xlsxHandler.Filename: xlsxFile,
	}

	for filename, file := range files {
		destination, err := os.OpenFile("./upload/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		io.Copy(destination, file)
	}

	defer ConverterJSONToXLSX()                        // call converter function
	http.Redirect(w, r, "/download", http.StatusFound) // redirect to download page don't remove this line
}

func ServeFileDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	http.ServeFile(w, r, "upload/report.xlsx")
}
