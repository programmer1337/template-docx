package utils

import "github.com/lukasjarosch/go-docx"

func PlaceholderReplacer(pathToTemplate string, pathToSave string, replaceMap docx.PlaceholderMap) {
	// read and parse the template docx
	// docx.ChangeOpenCloseDelimiter('$', '$')
	doc, err := docx.Open(pathToTemplate)
	if err != nil {
		panic(err)
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		panic(err)
	}

	err = doc.WriteToFile(pathToSave)
	if err != nil {
		panic(err)
	}
}
