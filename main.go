package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

type User struct {
	CPF  string `json:"CPF"`
	Data string `json:"Data"`
}

func main() {
	// Leitura do arquivo .json
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var users map[string]User

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Leitura do arquivo .xlsx
	excelFileName := "report.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// Acessando a coluna "Usuário"
			userXLSX := row.Cells[0].String()

			// Comparando o "CPF" em cada objeto do arquivo .json
			for _, userJSON := range users {
				if userJSON.CPF == userXLSX {
					fmt.Println("Match found:", userJSON.CPF)
				}
			}
		}
	}
}
