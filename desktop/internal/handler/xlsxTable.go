package handler

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// TODO
// - вернуть заголовки
// - полученные заголовки закинуть в именование полей формы

type XlsxTableHandler struct {
}

func NewXlsxTableHandler() *XlsxTableHandler {
	return &XlsxTableHandler{}
}

// TODO добавить возможность указывать с какой строки и с какого столбца начинать
func (xth *XlsxTableHandler) UpdateTable(file *excelize.File, sheet string, newData [][]any) error {
	rows, err := file.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cols, err := file.GetCols(sheet)
	if err != nil {
		fmt.Println(err)
		return err
	}

	xth.clearTable(file, sheet, []int{2, len(rows)}, []int{1, len(cols)})

	for row := 0; row < len(newData); row++ {
		for col := 0; col < len(newData[row]); col++ {
			// cell := fmt.Sprintf("%v%v", string('A'+rune(col)), row+2)
			// TODO добавить возможность указывать с какой строки и с какого столбца начинать
			cell, _ := excelize.CoordinatesToCellName(col+1, row+2)
			err := file.SetCellValue(sheet, cell, newData[row][col])
			if err != nil {
				log.Printf("Error setting cell %c%d: %v\n", col, row, err)
			}
		}
	}

	return nil
}

func (xth *XlsxTableHandler) GetRows(file *excelize.File, sheet string) ([][]string, error) {
	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (xth *XlsxTableHandler) clearTable(file *excelize.File, sheet string, rows []int, cols []int) {
	for row := rows[0]; row < rows[1]; row++ {
		for col := cols[0]; col < cols[1]; col++ {
			cell, _ := excelize.CoordinatesToCellName(col+1, row)
			err := file.SetCellValue(sheet, cell, "")
			if err != nil {
				log.Printf("Error setting cell %c%d: %v\n", col, row, err)
			}
		}
	}
}
