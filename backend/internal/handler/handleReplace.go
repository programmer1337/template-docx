package handler

import (
	"archive/zip"
	"bytes"
	entity "document-parser/internal/domain"
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

	counterparties := entity.Counterparties{}

	log.Print("ReadAll")
	body := new(bytes.Buffer)
	_, err := io.Copy(body, r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body.Bytes(), &counterparties)
	fmt.Printf("Received body: %v", counterparties)

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
		utils.PlaceholderReplacer(pathToTemplate, pathToSave, replaceMap)
	}

	downloadAllFiles(w)
}

func downloadAllFiles(w http.ResponseWriter) {
	dir := "../replaced"
	zipFileName := "all_files.zip"

	filesInDir, err := os.ReadDir(dir)
	if err != nil {
		http.Error(w, "Error! Can't create zip archive", http.StatusInternalServerError)
	}

	var filesToZip = FilesZipData{
		directory:     dir,
		filesDirEntry: filesInDir,
	}

	buf := new(bytes.Buffer)
	buf, err = createZipArchive(buf, filesToZip)
	if err != nil {
		http.Error(w, "Error! Can't create zip archive", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipFileName))
	w.WriteHeader(http.StatusOK)

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//--------------------------------//
//          ZIP ARCHIVE           //

type FilesZipData struct {
	directory     string
	filesDirEntry []os.DirEntry
}

func createZipArchive(buffer *bytes.Buffer, files FilesZipData) (*bytes.Buffer, error) {
	zipWriter := zip.NewWriter(buffer)
	defer zipWriter.Close()

	for _, file := range files.filesDirEntry {
		err := addFileToZip(zipWriter, filepath.Join(files.directory, file.Name()), file.Name())
		if err != nil {
			return nil, fmt.Errorf("Ошибка при добавлении файла в архив: %v", err)
		}
	}

	return buffer, nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла %s: %v", filePath, err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Ошибка при получении информации о файле %s: %v", filePath, err)
	}

	zipFileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return fmt.Errorf("Ошибка при создании заголовка файла %s: %v", filePath, err)
	}

	zipFileHeader.Name = fileName
	zipFileHeader.Method = zip.Deflate

	zipWriterEntry, err := zipWriter.CreateHeader(zipFileHeader)
	if err != nil {
		return fmt.Errorf("Ошибка при создании записи для файла %s: %v", filePath, err)
	}

	_, err = io.Copy(zipWriterEntry, file)
	if err != nil {
		return fmt.Errorf("Ошибка при копировании данных файла %s в архив: %v", filePath, err)
	}

	return nil
}
