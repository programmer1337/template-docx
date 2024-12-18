package handler

import (
	"archive/zip"
	"bytes"
	"document-parser/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gorilla/mux"
	"github.com/lukasjarosch/go-docx"
)

type Counteparty struct {
	Code_ou                               string `json:"code_ou"`
	Inn                                   string `json:"inn"`
	Institution_short_name                string `json:"institution_short_name"`
	Institution_full_name                 string `json:"institution_full_name"`
	Address                               string `json:"address"`
	City                                  string `json:"city"`
	Bank_details                          string `json:"bank_details"`
	Responsible_person_job_title          string `json:"responsible_person_job_title"`
	Responsible_person_short_name         string `json:"responsible_person_short_name"`
	Responsible_person_full_name          string `json:"responsible_person_full_name"`
	Responsible_person_full_name_genitive string `json:"responsible_person_full_name_genitive"`
	Acting_on                             string `json:"acting_on"`
	Ikz_2025                              string `json:"ikz_2025"`
	Source_funding                        string `json:"source_funding"`
	Email                                 string `json:"email"`
	Phone_number                          string `json:"phone_number"`
	Contract_form                         string `json:"contract_form"`
	Contract_type                         string `json:"contract_type"`
	Contract_number                       string `json:"contract_number"`
	Contract_formation_data               string `json:"contract_formation_data"`
	Responsible_person_job_title_genetive string `json:"responsible_person_job_title_genetive"`
	Category                              string `json:"category"`
}

type Counteparties []*Counteparty

func HandleReplace(serveMux *mux.Router, log *log.Logger) {
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/replace", Replace)
}

func Replace(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Expected Content-Type: application/json", http.StatusUnsupportedMediaType)
		return
	}

	log.Print("Decoding body")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	log.Printf("Received body (first 500 chars): %s", string(body[:min(500, len(body))]))
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var counteparties Counteparties
	err = json.NewDecoder(r.Body).Decode(&counteparties)
	if err != nil {
		log.Printf("Error decoding body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print("Body decoded successfully")

	var wg sync.WaitGroup
	wg.Add(len(counteparties))

	for _, counteparty := range counteparties {
		go func(counteparty *Counteparty) {
			defer wg.Done()

			replaceMap := docx.PlaceholderMap{
				"A": counteparty.Code_ou,
				"B": counteparty.Inn,
				"C": counteparty.Institution_short_name,
				"D": counteparty.Institution_full_name,
				"E": counteparty.Address,
				"F": counteparty.City,
				"G": counteparty.Bank_details,
				"H": counteparty.Responsible_person_job_title,
				"I": counteparty.Responsible_person_short_name,
				"J": counteparty.Responsible_person_full_name,
				"K": counteparty.Responsible_person_full_name_genitive,
				"L": counteparty.Acting_on,
				"M": counteparty.Ikz_2025,
				"N": counteparty.Source_funding,
				"O": counteparty.Email,
				"P": counteparty.Phone_number,
				"Q": counteparty.Contract_form,
				"R": counteparty.Contract_type,
				"S": counteparty.Contract_number,
				"T": counteparty.Contract_formation_data,
				"U": counteparty.Responsible_person_job_title_genetive,
				"V": counteparty.Category,
			}

			var pathToTemplate = "../templates/type" + counteparty.Contract_type + ".docx"
			var pathToSave = "../replaced/" + counteparty.Inn + ".docx"
			utils.PlaceholderReplacer(pathToTemplate, pathToSave, replaceMap)
		}(counteparty)
	}

	wg.Wait()

	downloadAllFiles(w)
}

func downloadAllFiles(w http.ResponseWriter) {
	log.Print("In downloadAllFiles method")

	dir := "../replaced"
	zipFileName := "all_files.zip"
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", zipFileName))

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Print("Error reading directory:", err)
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	// Используем sync.WaitGroup для параллельной обработки файлов
	var wg sync.WaitGroup
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".docx" {
			wg.Add(1) // Увеличиваем счетчик ожидания для каждого файла

			go func(file os.DirEntry) {
				defer wg.Done() // Уменьшаем счетчик ожидания по завершении горутины

				filePath := filepath.Join(dir, file.Name())
				f, err := os.Open(filePath)
				if err != nil {
					log.Printf("Unable to open file %s: %v", file.Name(), err)
					return
				}
				defer f.Close()

				zipFile, err := zipWriter.Create(file.Name())
				if err != nil {
					log.Printf("Unable to create zip entry for %s: %v", file.Name(), err)
					return
				}

				_, err = io.Copy(zipFile, f)
				if err != nil {
					log.Printf("Error copying file %s data: %v", file.Name(), err)
				}
			}(file)
		}
	}

	wg.Wait()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
