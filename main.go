package main

import (
	"conversor/internal"
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/upload", internal.UploadHandler)
	http.HandleFunc("/", internal.HtmlHandler)
	http.HandleFunc("/download", internal.DownloadHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
