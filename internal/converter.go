package internal

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

const (
	JsonFileName = "./upload/data.json"
	XlsxFileName = "./upload/data.xlsx"
)

func ConverterJSONToXLSX() error {
	jsFile, err := os.Open(JsonFileName)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo JSON: %v", err)
	}
	defer jsFile.Close()

	var users map[string]User
	err = json.NewDecoder(jsFile).Decode(&users)
	if err != nil {
		return fmt.Errorf("erro ao decodificar o arquivo JSON: %v", err)
	}

	xlFile, err := xlsx.OpenFile(XlsxFileName)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo XLSX: %v", err)
	}

	var matches [][]string
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) > 1 {
				userXLSX := row.Cells[0].String()
				nameXLSX := row.Cells[1].String()

				for _, userJSON := range users {
					if userJSON.CPF == userXLSX {
						matches = append(matches, []string{userJSON.CPF, nameXLSX})
					}
				}
			}
		}
	}
	saveXLSXFile(matches)

	return nil
}

func saveXLSXFile(matches [][]string) error {
	fileX := xlsx.NewFile()
	sheet, err := fileX.AddSheet("Matches")
	if err != nil {
		return fmt.Errorf("erro ao criar uma nova planilha: %v", err)
	}

	for _, match := range matches {
		row := sheet.AddRow()
		row.AddCell().Value = match[0]
		row.AddCell().Value = match[1]
	}

	err = fileX.Save("./upload/report.xlsx")
	if err != nil {
		return fmt.Errorf("erro ao salvar o arquivo XLSX: %v", err)
	}

	return nil
}
