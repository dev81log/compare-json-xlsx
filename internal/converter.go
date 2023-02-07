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

func LogicConverterFiles() {
	file, err := os.Open("./upload/data.json")
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

	excelFileName := "./upload/data.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		return
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
	fileX := xlsx.NewFile()
	sheet, err := fileX.AddSheet("Matchs")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, match := range matches {
		row := sheet.AddRow()
		row.AddCell().Value = match[0]
		row.AddCell().Value = match[1]
	}

	err = fileX.Save("./upload/relatorio.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matchs salvo em relatorio.xlsx")
}
