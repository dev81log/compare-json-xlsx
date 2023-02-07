package main

import (
	"conversor/internal"
	"fmt"
	"net/http"
	"text/template"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/fileUpload.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	jsonFile, _, err := r.FormFile("jsonFile")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo json: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer jsonFile.Close()

	xlsxFile, _, err := r.FormFile("xlsxFile")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo xlsx: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer xlsxFile.Close()

	internal.LogicConverterFiles()
	w.Write([]byte("Arquivos convertidos com sucesso!"))

}

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", htmlHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
