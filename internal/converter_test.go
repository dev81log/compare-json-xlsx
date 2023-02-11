package internal

import (
	"os"
	"testing"
)

func TestConverterJSONToXLSX(t *testing.T) {

	jsFile, err := os.Create("." + JsonFileName)
	if err != nil {
		t.Fatalf("erro ao criar o arquivo JSON: %v", err)
	}
	defer jsFile.Close()

	_, err = jsFile.Write([]byte(`{"123456": {"CPF": "123456", "Data": "2022-01-01"}}`))
	if err != nil {
		t.Fatalf("erro ao escrever no arquivo JSON: %v", err)
	}

	xlFile, err := os.Create("." + XlsxFileName)
	if err != nil {
		t.Fatalf("erro ao criar o arquivo XLSX: %v", err)
	}
	defer xlFile.Close()

	_, err = os.Stat("../upload/report.xlsx")
	if err != nil {
		t.Fatalf("erro ao verificar se o arquivo report.xlsx foi criado: %v", err)
	}

}
