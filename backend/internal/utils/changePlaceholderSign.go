package utils

import (
	"fmt"

	"github.com/nguyenthenguyen/docx"
)

func changeTemplateSign(pathToTemplate string, pathToSave string) {
	dataMap := map[string]string{
		"} {": "} {",
	}

	r, err := docx.ReadDocxFile(pathToTemplate)

	if err != nil {
		panic(err)
	}
	loadedDocx := r.Editable()

	for key, value := range dataMap {
		err = loadedDocx.Replace(key, value, -1)
		if err != nil {
			fmt.Printf("Ошибка при замене '%s': %v\n", key, err)
		}
	}
	loadedDocx.WriteToFile(pathToSave)

	r.Close()
}
