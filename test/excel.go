package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	//excel 测试~
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	//f.SetCellValue("Sheet2", "B2", 100)
	for i:=1;i<10;i++{
		s:=strconv.Itoa(i)
		f.SetCellValue("Sheet2", "A"+s, i)
		f.SetCellValue("Sheet2", "B"+s, i)
		f.SetCellValue("Sheet2", "C"+s, i)

	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}