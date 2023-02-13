![Asher_Duarte_icon_styles_machine_converter_data_155cef66-7b0d-46bc-982f-681cef56daad (1)](https://user-images.githubusercontent.com/105469529/218542780-2f4b7bec-d4bc-4476-8b55-6ebf05cc3bc2.png)
# Compare-json-xlsx converter
This project is a tool for converting a JSON file into an XLSX spreadsheet. The tool compares the contents of the JSON file with the contents of an XLSX spreadsheet and creates a new XLSX file with the matching data.

## Technologies Used
+ Go
+ Library for manipulating XLSX spreadsheets - ([github.com/tealeg/xlsx](url)) 

## How it Works
`The tool starts by opening the JSON file and decoding its contents into a Go data structure. Next, it opens the XLSX spreadsheet and iterates over its cells, comparing the contents with the contents of the JSON file. When there is a match, the data is added to a match array. Finally, a new XLSX file is created with the match array.`

## Usage
To use the tool, simply call the ConverterJSONToXLSX() function. If there is an error during the process, an error message will be displayed. The new XLSX file will be saved in the `"./upload/report.xlsx"` directory.

### Final Considerations
This project can be used as a basis for building more complex tools for manipulating data in JSON and XLSX files. Feel free to contribute and improve it!
