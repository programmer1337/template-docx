package handler

import (
	entity "document-parser/internal/domain"
	"document-parser/internal/utils"
	"document-parser/pkg/ziputils"
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

func HandleReplace(serveMux *mux.Router, log *log.Logger) {
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/replace", Replace)
}

// Функция для замены значения поля структуры по имени
func Replace(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		http.Error(w, "Expected Content-Type: application/json", http.StatusUnsupportedMediaType)
		return
	}

	counterparties := entity.Counterparties{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &counterparties)

	err = os.RemoveAll("../replaced/")
	if err != nil {
		log.Println(err)
		return
	}

	for _, conteragent := range counterparties {
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

		var pathToTemplate = "../templates/type" + conteragent.Contract_type + ".docx"
		var pathToSave = "../replaced/" + conteragent.Inn + ".docx"

		err = utils.PlaceholderReplacer(pathToTemplate, pathToSave, replaceMap)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error with placeholder [%v]", err), http.StatusInternalServerError)
			return
		}
	}

	//TODO Исправить
	downloadAllFiles(w)
}

func downloadAllFiles(w http.ResponseWriter) {
	directory := "../replaced"
	zipFileName := "all_files.zip"

	//TODO формировать files с помощью бд
	files, err := getFiles(directory)
	if err != nil {
		http.Error(w, "Problem's with file's", http.StatusInternalServerError)
		return
	}

	buf, err := ziputils.CreateZipArchive(ziputils.FilesZipData{
		Directory: directory,
		Files:     files,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error! Can't create zip archive. [%v]", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipFileName))
	w.WriteHeader(http.StatusOK)

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't add to buffer: %v", err.Error()), http.StatusInternalServerError)
		return
	}
}

func getFiles(directory string) ([]ziputils.File, error) {
	filesDirEntry, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("Error! Can't create zip archive")
	}

	files := []ziputils.File{}
	for _, file := range filesDirEntry {
		files = append(files, ziputils.File{
			Name:      file.Name(),
			Directory: filepath.Join(directory, file.Name()),
		})
	}

	return files, nil
}
