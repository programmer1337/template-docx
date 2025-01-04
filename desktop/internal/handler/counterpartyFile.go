package handler

import (
	"desktop-templater-docx/internal/domain/entity"
	"desktop-templater-docx/pkg/structutils"
	"fmt"
	"io"
	"log"

	"github.com/xuri/excelize/v2"
)

// TODO
// - вернуть заголовки
// - полученные заголовки закинуть в именование полей формы
type CounterpartyFileHandler struct {
	file *excelize.File
}

func NewCounterpartyFileHandler() *CounterpartyFileHandler {
	return &CounterpartyFileHandler{
		file: nil,
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
	rows, err := cfh.file.GetRows(cfh.file.GetSheetList()[0])
	if err != nil {
		return nil, err
	}

	counterparties := entity.Counterparties{}

	if len(rows) > 1 {
		headers := rows[0]

		for _, row := range rows[1:] {
			//TODO подумать
			counterparty := entity.Counterparty{
				Code_ou:                               "",
				Inn:                                   "",
				Institution_short_name:                "",
				Institution_full_name:                 "",
				Address:                               "",
				City:                                  "",
				Bank_details:                          "",
				Responsible_person_job_title:          "",
				Responsible_person_short_name:         "",
				Responsible_person_full_name:          "",
				Responsible_person_full_name_genitive: "",
				Acting_on:                             "",
				Ikz_2025:                              "",
				Source_funding:                        "",
				Email:                                 "",
				Phone_number:                          "",
				Contract_form:                         "",
				Contract_type:                         "",
				Contract_number:                       "",
				Contract_formation_data:               "",
				Responsible_person_job_title_genetive: "",
				Category:                              "",
			}

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

	for row := 0; row <= len(counterparties)-1; row++ {
		counterparty := structutils.GetStructValues(*counterparties[row])

		for col := 0; col <= len(counterparty)-1; col++ {
			cell := fmt.Sprintf("%v%v", string('A'+rune(col)), row+2)

			err := cfh.file.SetCellValue(sheet, cell, counterparty[col])
			if err != nil {
				log.Printf("Error setting cell %c%d: %v\n", col, row, err)
			}
		}
	}

	// fmt.Println(cfh.file.Path)
	cfh.file.Save()
}
