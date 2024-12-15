package utils

import "github.com/lukasjarosch/go-docx"

func PlaceholderReplacer(pathToTemplate string, pathToSave string, replaceMap docx.PlaceholderMap) {
	// replaceMap := docx.PlaceholderMap{
	// 	"INN":                   "2301033998",
	// 	"INSTITUTION_FULL_NAME": "Наименование учебного заведения",
	// }

	// read and parse the template docx
	// docx.ChangeOpenCloseDelimiter('$', '$')
	doc, err := docx.Open(pathToTemplate)
	if err != nil {
		panic(err)
	}

	// replace the keys with values from replaceMap
	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		panic(err)
	}

	err = doc.WriteToFile(pathToSave)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	dataMap := map[string]string{
// 		"{INN}":                   "2301033998",
// 		"{INSTITUTION_FULL_NAME}": "Наименование учебного заведения",
// 	}

// 	// Read from docx file
// 	r, err := docx.ReadDocxFile("./type1.docx")

// 	if err != nil {
// 		panic(err)
// 	}
// 	loadedDocx := r.Editable()

// 	content := loadedDocx.GetContent()
// 	log.Print(content)

// 	for key, value := range dataMap {
// 		err = loadedDocx.Replace(key, value, -1)
// 		if err != nil {
// 			fmt.Printf("Ошибка при замене '%s': %v\n", key, err)
// 		}
// 	}
// 	loadedDocx.WriteToFile("./new_result_1.docx")

// 	r.Close()
// }
