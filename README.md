# Compare-json-xlsx converter
This project is a tool for converting a JSON file into an XLSX spreadsheet. The tool compares the contents of the JSON file with the contents of an XLSX spreadsheet and creates a new XLSX file with the matching data.

## Technologies Used
```
Go
github.com/tealeg/xlsx - Library for manipulating XLSX spreadsheets
```
How it Works
The tool starts by opening the JSON file and decoding its contents into a Go data structure. Next, it opens the XLSX spreadsheet and iterates over its cells, comparing the contents with the contents of the JSON file. When there is a match, the data is added to a match array. Finally, a new XLSX file is created with the match array.

## Usage
To use the tool, simply call the ConverterJSONToXLSX() function. If there is an error during the process, an error message will be displayed. The new XLSX file will be saved in the "./upload/report.xlsx" directory.

### Final Considerations
This project can be used as a basis for building more complex tools for manipulating data in JSON and XLSX files. Feel free to contribute and improve it!
