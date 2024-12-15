package utils

import (
	"fmt"

	"github.com/nguyenthenguyen/docx"
)

func changeTemplateSign(pathToTemplate string, pathToSave string) {
	// dataMap := map[string]string{
	// 	"{U}":                     "директора",
	// 	"{INSTITUTION_FULL_NAME}": "Наименование учебного заведения",
	// }

	dataMap := map[string]string{
		"} {": "} {",
	}

	// Read from docx file
	r, err := docx.ReadDocxFile(pathToTemplate)

	if err != nil {
		panic(err)
	}
	loadedDocx := r.Editable()

	// content := loadedDocx.GetContent()
	// log.Print(content)

	for key, value := range dataMap {
		err = loadedDocx.Replace(key, value, -1)
		if err != nil {
			fmt.Printf("Ошибка при замене '%s': %v\n", key, err)
		}
	}
	loadedDocx.WriteToFile(pathToSave)

	r.Close()
}
