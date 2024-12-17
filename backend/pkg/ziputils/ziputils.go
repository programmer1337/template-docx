package ziputils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

// --------------------------------//
type FilesZipData struct {
	Directory string
	Files     []File
}

type File struct {
	Name      string
	Directory string
}

func CreateZipArchive(files FilesZipData) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buffer)
	defer zipWriter.Close()

	for _, file := range files.Files {
		err := addFileToZip(zipWriter, file.Directory, file.Name)
		if err != nil {
			return nil, fmt.Errorf("Error file add [%v]", err)
		}
	}

	return buffer, nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Can't open file. Error %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Can't get file info. Error %v", err)
	}

	zipFileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return fmt.Errorf("Can't create file header. Error %v", err)
	}

	zipFileHeader.Name = fileName
	zipFileHeader.Method = zip.Deflate

	zipWriterEntry, err := zipWriter.CreateHeader(zipFileHeader)
	if err != nil {
		return fmt.Errorf("Can't create writer. Error %v", err)
	}

	_, err = io.Copy(zipWriterEntry, file)
	if err != nil {
		return fmt.Errorf("Can't copy to archive. Error %v", err)
	}

	return nil
}
