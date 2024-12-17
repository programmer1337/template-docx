package handler

import (
	"archive/zip"
	"document-parser/internal/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"github.com/lukasjarosch/go-docx"
)

type Counterparty struct {
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

type Counterparties []*Counterparty

type KeyCounterparties struct{}

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

	// ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	// defer cancel()

	counterparties := Counterparties{}

	log.Print("ReadAll")
	body, err := io.ReadAll(r.Body)
	r.Body.Close()

	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&counterparties)

	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	// err = json.Unmarshal(body, &counterparties)
	// if err != nil {
	// 	fmt.Println(w, "can't unmarshal: ", err.Error())
	// }

	// if err != nil {
	// 	if ctx.Err() == context.DeadlineExceeded {
	// 		log.Println("Request body read timed out")
	// 	}
	// 	log.Printf("Error reading request body: %v", err)
	// 	http.Error(w, "Unable to read request body", http.StatusInternalServerError)
	// 	return
	// }

	log.Print("jsoniter")
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &counterparties)

	fmt.Printf("Received body: %v", counterparties)

	log.Print("RemoveAll")
	err = os.RemoveAll("../replaced/")
	if err != nil {
		log.Println(err)
		return
	}

	log.Print("counterparties")
	for _, conteragent := range counterparties {
		// log.Println(pos, conteragent)

		replaceMap := docx.PlaceholderMap{
			"A": conteragent.Code_ou,
			"B": conteragent.Inn,
			"C": conteragent.Institution_short_name,
			"D": conteragent.Institution_full_name,
			"E": conteragent.Address,
			"F": conteragent.City,
			"G": conteragent.Bank_details,
			"H": conteragent.Responsible_person_job_title,
			"I": conteragent.Responsible_person_short_name,
			"J": conteragent.Responsible_person_full_name,
			"K": conteragent.Responsible_person_full_name_genitive,
			"L": conteragent.Acting_on,
			"M": conteragent.Ikz_2025,
			"N": conteragent.Source_funding,
			"O": conteragent.Email,
			"P": conteragent.Phone_number,
			"Q": conteragent.Contract_form,
			"R": conteragent.Contract_type,
			"S": conteragent.Contract_number,
			"T": conteragent.Contract_formation_data,
			"U": conteragent.Responsible_person_job_title_genetive,
			"V": conteragent.Category,
		}

		// "./templates/type1.docx"
		var pathToTemplate = "../templates/type" + conteragent.Contract_type + ".docx"
		var pathToSave = "../replaced/" + conteragent.Inn + ".docx"
		utils.PlaceholderReplacer(pathToTemplate, pathToSave, replaceMap)
	}
	downloadAllFiles(w)
}

func downloadAllFiles(w http.ResponseWriter) {
	// Путь к папке с файлами
	dir := "../replaced"

	// Создаем новый архив
	zipFileName := "all_files.zip"
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", zipFileName))

	// Создаем новый zip.Writer
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// Получаем список всех файлов в директории
	files, err := os.ReadDir(dir)
	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	// Перебираем все файлы и добавляем их в архив
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".docx" {
			// Открываем файл для чтения
			filePath := filepath.Join(dir, file.Name())
			f, err := os.Open(filePath)
			if err != nil {
				http.Error(w, "Unable to open file", http.StatusInternalServerError)
				return
			}
			defer f.Close()

			// Добавляем файл в архив
			zipFile, err := zipWriter.Create(file.Name())
			if err != nil {
				http.Error(w, "Unable to create zip entry", http.StatusInternalServerError)
				return
			}

			// Копируем содержимое файла в архив
			_, err = io.Copy(zipFile, f)
			if err != nil {
				http.Error(w, "Error copying file data", http.StatusInternalServerError)
				return
			}
		}
	}
}
