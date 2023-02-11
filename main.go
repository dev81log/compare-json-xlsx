package main

import (
	"conversor/internal"
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", internal.RenderHTML)
	http.HandleFunc("/upload", internal.HandleFileUpload)
	http.HandleFunc("/download", internal.ServeFileDownload)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
