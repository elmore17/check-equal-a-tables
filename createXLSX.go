package main

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

func createSheetXlsx(f *excelize.File, nameSheet string, columnName []string, param1 string, param2 string) *excelize.File {
	sheetName := nameSheet

	index, err := f.NewSheet(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	f.SetActiveSheet(index)

	for col, header := range columnName {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	rowNum := 2
	f.SetCellValue(sheetName, "A"+strconv.Itoa(rowNum), param1)
	f.SetCellValue(sheetName, "B"+strconv.Itoa(rowNum), param2)

	return f
}
