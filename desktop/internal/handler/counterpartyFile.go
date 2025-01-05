package handler

import (
	"desktop-templater-docx/internal/domain/entity"
	"desktop-templater-docx/pkg/structutils"
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

type CounterpartyFileHandler struct {
	file         *excelize.File
	tableHandler *XlsxTableHandler
}

func NewCounterpartyFileHandler() *CounterpartyFileHandler {
	return &CounterpartyFileHandler{
		file:         nil,
		tableHandler: NewXlsxTableHandler(),
	}
}

func (cfh *CounterpartyFileHandler) SetCounterpartiesFileByReader(reader io.Reader, path string) error {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return fmt.Errorf("have problems %v", err)
	}
	file.Path = path
	cfh.file = file

	return nil
}

func (cfh *CounterpartyFileHandler) LoadCounterparties() (entity.Counterparties, error) {
	rows, err := cfh.tableHandler.GetRows(cfh.file, cfh.file.GetSheetList()[0])
	if err != nil {
		return nil, err
	}

	counterparties := entity.Counterparties{}

	if len(rows) > 1 {
		headers := rows[0]

		for _, row := range rows[1:] {
			var counterparty entity.Counterparty

			for i, value := range row {
				err = structutils.SetFieldValue(&counterparty, entity.CounterpartyRuEngMap[headers[i]], value)
				if err != nil {
					return nil, err
				}
			}

			counterparties = append(counterparties, &counterparty)
		}
	}

	return counterparties, nil
}

func (cfh *CounterpartyFileHandler) SaveCounterparties(counterparties entity.Counterparties) {
	if cfh.file == nil {
		return
	}

	sheet := cfh.file.GetSheetList()[0]

	formattedData := [][]any{}

	for row := 0; row <= len(counterparties)-1; row++ {
		counterparty := structutils.GetStructValues(*counterparties[row])

		formattedData = append(formattedData, counterparty)
	}

	// fmt.Println(cfh.file.Path)
	cfh.tableHandler.UpdateTable(cfh.file, sheet, formattedData)
	cfh.file.Save()
}
