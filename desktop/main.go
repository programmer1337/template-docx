package main

import (
	"desktop-templater-docx/internal/creator"
	"desktop-templater-docx/internal/domain/entity"
	"desktop-templater-docx/internal/handler"
	"desktop-templater-docx/pkg/customui"
	"desktop-templater-docx/pkg/structutils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("templater-docx")
	w := app.NewWindow("Работа с контрагентами")

	cfh := handler.NewCounterpartyFileHandler()
	ch := handler.NewCounterpartyHandler()

	ListTitleLabels := customui.NewListTitleLabels(
		func(pos int, item *customui.ListItemTitleLabels) {
			counterparty, err := ch.GetCounterparty(pos)
			if err != nil {
				fmt.Printf("[Error]: %v\n", err)
				return
			}
			openCounterparties(app, counterparty, item)
		},
		nil,
		// func(isSelected bool) {
		// 	if !isSelected {
		// 		selectAllCheck.
		// 	}
		// },
	)

	//--------------------------------------------//
	// UI
	loadButton := widget.NewButton("Загрузить файл", func() {
		fileDialog := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, err error) {
				if err != nil {
					fmt.Printf("[Error]: while opening: %v\n", err)
					return
				}
				if uc != nil {
					err = cfh.SetCounterpartiesFileByReader(uc, uc.URI().Path())
					if err != nil {
						fmt.Printf("[Error]: %v\n", err)
					}
					// TODO
					counterparties, err := cfh.LoadCounterparties()
					ch.SetCounterparties(counterparties)
					if err != nil {
						fmt.Printf("[Error]: %v\n", err)
					}

					counterparties = ch.GetCounterparties()
					ListTitleLabels.SetItems(counterparties)
					uc.Close()
				} else {
					fmt.Println("Файл не выбран")
				}
			}, w)

		fileDialog.Show()
	})
	saveButton := widget.NewButton("Сохранить файл", func() {
		counterparties := ch.GetCounterparties()
		cfh.SaveCounterparties(counterparties)
	})
	// 	if changed {
	// 		ListTitleLabels.SelectAll()
	// 	} else {
	// 		ListTitleLabels.UnselectAll()
	// 	}
	// })
	processTemplateButton := widget.NewButton("Сформировать договора", func() {
		counterpartiesToProcess := ListTitleLabels.GetSelected()

		for _, id := range counterpartiesToProcess {
			ch.MakeReplace(id)
		}
	})
	//--------------------------------------------//

	content := container.NewBorder(
		// container.NewVBox(container.NewHBox(loadButton, saveButton), container.NewHBox(selectAllCheck)),
		container.NewVBox(container.NewHBox(loadButton, saveButton)),
		container.NewVBox(processTemplateButton),
		nil,
		nil,
		// container.NewScroll(listContainer),
		container.NewScroll(ListTitleLabels),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1280, 720))
	w.ShowAndRun()
}

func openCounterparties(app fyne.App, counterparty *entity.Counterparty, litl *customui.ListItemTitleLabels) {
	ew := app.NewWindow("Редактирование контрагента")

	onSave := func(inputs []creator.InputPlaceholder) {
		for _, item := range inputs {
			structutils.SetFieldValue(counterparty, item.Name, item.Input.Text)
		}
		litl.Title.SetText(counterparty.Inn)
		litl.SetLabels(
			[]string{
				counterparty.Institution_short_name,
				counterparty.Responsible_person_short_name,
				counterparty.City,
			})
		ew.Close()
	}
	onCancel := func() {
		ew.Close()
	}

	formContent := creator.CreateForm(onSave, onCancel, counterparty)

	ew.SetContent(formContent)
	ew.Resize(fyne.NewSize(1280, 720))
	ew.Show()
}

/*func openCounterparties(window fyne.Window, counterparty *entity.Counterparty) {
	inputContainer := container.NewVBox()

	fields := structutils.GetStructFieldNames(entity.Counterparty{})
	inputs := []InputPlaceholder{}
	for _, fieldName := range fields {
		input := widget.NewEntry()
		if fieldName == "Bank_details" {
			input.SetMinRowsVisible(10)
		}
		input.MultiLine = true
		input.SetPlaceHolder(fmt.Sprintf("Введите %v", entity.CounterpartyEngRuAlias[fieldName]))
		currentValue, err := structutils.GetFieldValue(counterparty, fieldName)
		if err != nil {
			fmt.Printf("Error: %v", err)
			// return
		}
		input.SetText(currentValue.String())
		inputs = append(inputs, InputPlaceholder{input: input, name: fieldName, placeholder: entity.CounterpartyEngRuAlias[fieldName]})

		inputContainer.Add(widget.NewLabel(fmt.Sprintf("%v:", entity.CounterpartyEngRuAlias[fieldName])))
		inputContainer.Add(input)
	}

	//----------Buttons----------//
	saveButton := widget.NewButton("Сохранить", func() {
		for _, item := range inputs {
			structutils.SetFieldValue(counterparty, item.name, item.input.Text)
		}
		window.Close()
		// customList[3].Title.SetText(counterparty.Inn)
	})

	cancelButton := widget.NewButton("Отменить", func() {
		window.Close()
	})

	buttonContainer := container.NewHBox(saveButton, cancelButton)
	//---------------------------//

	editContent := container.NewBorder(
		widget.NewLabel(fmt.Sprintf("Редактирование контрагента %v", counterparty.Inn)),
		nil,
		buttonContainer,
		nil,
		nil,
		container.NewScroll(container.NewGridWithColumns(2, inputContainer)),
	)

	// return editContent

	window.SetContent(editContent)
	window.Resize(fyne.NewSize(1280, 720))
	window.Show()
}*/
