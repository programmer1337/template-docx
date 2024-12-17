package handler

import (
	entity "document-parser/internal/domain"
	"document-parser/pkg/structutils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
)

func HandleCounterparty(serveMux *mux.Router, log *log.Logger) {
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/loadCounterpartiesXlsx", LoadXlsx)
}

//TODO получить страницу, которую нужно обработать и файл
//TODO вынести работу с контрагентом, добавить ToJSON и FromJSON, работу с контекстом

// type XlsxRequest struct {
// 	CurrentSheet int
// 	XlsxFile     []byte
// }

func LoadXlsx(w http.ResponseWriter, r *http.Request) {
	file, err := excelize.OpenReader(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err.Error()), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	defer func() {
		if err := file.Close(); err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err.Error()), http.StatusInternalServerError)
		}
	}()

	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		http.Error(w, fmt.Sprintf("Excel file have problem: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	counterparties := entity.Counterparties{}

	if len(rows) > 0 {
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
				error := structutils.SetFieldValue(&counterparty, entity.CounterpartyAlias[headers[i]], value)
				if error != nil {
					fmt.Println("Ошибка:", error)
					return
				}
			}

			counterparties = append(counterparties, &counterparty)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(counterparties); err != nil {
		http.Error(w, fmt.Sprintf("Server encode error: %v", err.Error()), http.StatusInternalServerError)
		return
	}
}
