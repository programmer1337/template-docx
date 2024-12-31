package utils

import (
	"fmt"

	"github.com/lukasjarosch/go-docx"
)

func PlaceholderReplacer(pathToTemplate string, pathToSave string, replaceMap docx.PlaceholderMap) error {
	// read and parse the template docx
	// docx.ChangeOpenCloseDelimiter('$', '$')
	doc, err := docx.Open(pathToTemplate)
	if err != nil {
		return fmt.Errorf("Can't open file")
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return fmt.Errorf("Can't replace. [%v]", err)
	}

	err = doc.WriteToFile(pathToSave)
	if err != nil {
		return fmt.Errorf("Can't write. [%v]", err)
	}

	return nil
}
