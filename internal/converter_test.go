package internal

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/tealeg/xlsx"
)

func TestConverterFiles(t *testing.T) {
	type User struct {
		CPF  string `json:"CPF"`
		Data string `json:"Data"`
	}
	users := map[string]User{
		"12345678901": User{CPF: "12345678901", Data: "01/01/2020"},
		"23456789012": User{CPF: "23456789012", Data: "02/02/2020"},
		"34567890123": User{CPF: "34567890123", Data: "03/03/2020"},
	}
	file, err := os.Create("../upload/data.json")
	if err != nil {
		t.Errorf("Erro ao criar arquivo JSON: %v", err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(users)

	excelFile := xlsx.NewFile()
	sheet, err := excelFile.AddSheet("Dados")
	if err != nil {
		t.Errorf("Erro ao adicionar planilha ao arquivo XLSX: %v", err)
	}

	for _, user := range users {
		row := sheet.AddRow()
		row.AddCell().Value = user.CPF
		row.AddCell().Value = "Nome" + user.CPF
	}

	err = excelFile.Save("../upload/data.xlsx")
	if err != nil {
		t.Errorf("Erro ao salvar arquivo XLSX: %v", err)
	}

	ConverterFiles()

	relatorio, err := xlsx.OpenFile("../upload/relatorio.xlsx")
	if err != nil {
		t.Errorf("Erro ao abrir arquivo de relatório: %v", err)
	}

	if len(relatorio.Sheets) != 1 {
		t.Errorf("Arquivo de relatório deve ter apenas uma planilha, mas possui %d", len(relatorio.Sheets))
	}

	sheet = relatorio.Sheets[0]
	if sheet.Name != "Matches" {
		t.Errorf("Planilha deve ter o nome 'Matches', mas tem o nome '%s'", sheet.Name)
	}

	if _, err := os.Stat("../upload/data.json"); os.IsNotExist(err) {
		t.Errorf("Arquivo não encontrado: %s", err)
	}
	if _, err := os.Stat("../upload/data.xlsx"); os.IsNotExist(err) {
		t.Errorf("Arquivo não encontrado: %s", err)
	}
	if _, err := os.Stat("../upload/relatorio.xlsx"); os.IsNotExist(err) {
		t.Errorf("Arquivo não encontrado: %s", err)
	}
}
