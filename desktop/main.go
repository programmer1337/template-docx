package main

import (
	entity "desktop-templater-docx/internal/domain"
	"desktop-templater-docx/pkg/customui"
	"desktop-templater-docx/pkg/sliceutils"
	"desktop-templater-docx/pkg/structutils"
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

type InputPlaceholder struct {
	input       *widget.Entry
	name        string
	placeholder string
}

func main() {
	app := app.New()
	w := app.NewWindow("Работа с контрагентами")

	counterpartiesList := entity.Counterparties{}
	customListItems := []*customui.MyListItemWidget{}
	listContainer := container.NewVBox()

	selectedItems := []*int{}

	// list.OnSelected = func(id int) {
	// 	fmt.Println(id)

	// 	editWindow := app.NewWindow("Редактирование контрагента")

	// 	inputContainer := container.NewVBox()

	// 	fields := structutils.GetStructFieldNames(entity.Counterparty{})
	// 	inputs := []InputPlaceholder{}
	// 	for _, fieldName := range fields {
	// 		input := widget.NewEntry()
	// 		if fieldName == "Bank_details" {
	// 			input.MultiLine = true
	// 			input.SetMinRowsVisible(10)
	// 		}
	// 		input.SetPlaceHolder(fmt.Sprintf("Введите %v", entity.CounterpartyEngRuAlias[fieldName]))
	// 		currentValue, err := structutils.GetFieldValue(counterpartiesList[id], fieldName)
	// 		if err != nil {
	// 			fmt.Printf("Error: %v", err)
	// 			return
	// 		}
	// 		input.SetText(currentValue.String())
	// 		inputs = append(inputs, InputPlaceholder{input: input, name: fieldName, placeholder: entity.CounterpartyEngRuAlias[fieldName]})

	// 		inputContainer.Add(widget.NewLabel(fmt.Sprintf("%v:", entity.CounterpartyEngRuAlias[fieldName])))
	// 		inputContainer.Add(input)
	// 	}

	// 	//----------Buttons----------//
	// 	saveButton := widget.NewButton("Сохранить", func() {
	// 		for _, item := range inputs {
	// 			structutils.SetFieldValue(counterpartiesList[id], item.name, item.input.Text)
	// 		}
	// 		// items[id] = entry.Text
	// 		editWindow.Close()
	// 		list.Unselect(id)
	// 		list.Refresh()
	// 		// fmt.Printf("Item edited: %s\n", input.Text)
	// 	})

	// 	cancelButton := widget.NewButton("Отменить", func() {
	// 		list.Unselect(id)
	// 		editWindow.Close()
	// 	})

	// 	buttonContainer := container.NewHBox(saveButton, cancelButton)
	// 	//---------------------------//

	// 	editContent := container.NewBorder(
	// 		// widget.NewLabel("Edit Item:"),
	// 		nil,
	// 		buttonContainer,
	// 		nil,
	// 		nil,
	// 		container.NewScroll(inputContainer),
	// 	)

	// 	editWindow.Resize(fyne.NewSize(1280, 720))
	// 	editWindow.SetContent(editContent)
	// 	editWindow.Show()
	// }

	selectAllButton := widget.NewButton("Выбрать все", func() {
		for i := range len(customListItems) {
			customListItems[i].Check.SetChecked(true)
		}
	})
	loadButton := widget.NewButton("Загрузить файл", func() {
		fileDialog := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, err error) {
				if err != nil {
					fmt.Printf("Error while opening: %v", err)
					return
				}
				if uc != nil {
					counterparties, err := loadCounterparties(uc)
					if err != nil {
						fmt.Printf("Programm stop with error [%v]", err)
					}

					fmt.Printf("Выбран файл: %s\n", uc.URI().Path())
					counterpartiesList = append(counterpartiesList, *counterparties...)
					for pos, c := range counterpartiesList {
						comment := c.Institution_short_name + " | " + c.Responsible_person_short_name + " | " + c.City
						customListItems = append(customListItems, customui.NewMyListItemWidget(
							c.Inn,
							comment,
							func(changed bool) {
								if changed {
									selectedItems = append(selectedItems, &pos)
								} else {
									selectedItems = sliceutils.RemoveByValue(selectedItems, &pos)
								}
								fmt.Printf("%v was changed: %v\n", pos, changed)
								fmt.Printf("selected items was changed: %v\n", selectedItems)
							},
							func() {
								fmt.Printf("Was tapped: %v\n", pos)
							}))
					}

					for _, item := range customListItems {
						listContainer.Add(item)
					}

					// listContainer.Refresh()
					uc.Close()
				} else {
					fmt.Println("Файл не выбран")
				}
			}, w)

		fileDialog.Show()
	})

	processTemplateButton := widget.NewButton("Сформировать договора", func() {})
	content := container.NewBorder(container.NewVBox(selectAllButton, loadButton), container.NewVBox(processTemplateButton), nil, nil, container.NewScroll(listContainer))

	w.SetContent(content)

	w.Resize(fyne.NewSize(1280, 720))
	w.ShowAndRun()
}

func loadCounterparties(reader io.Reader) (*entity.Counterparties, error) {
	file, err := excelize.OpenReader(reader)

	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		return nil, err
	}

	counterparties := entity.Counterparties{}

	if len(rows) > 0 {
		headers := rows[0]

		for _, row := range rows[1:] {
			//TODO подумать
			counterparty := entity.Counterparty{
				Code_ou:                               "",
				Inn:                                   "",
				Institution_short_name:                "",
				Institution_full_name:                 "",
				Address:                               "",
				City:                                  "",
				Bank_details:                          "",
				Responsible_person_job_title:          "",
				Responsible_person_short_name:         "",
				Responsible_person_full_name:          "",
				Responsible_person_full_name_genitive: "",
				Acting_on:                             "",
				Ikz_2025:                              "",
				Source_funding:                        "",
				Email:                                 "",
				Phone_number:                          "",
				Contract_form:                         "",
				Contract_type:                         "",
				Contract_number:                       "",
				Contract_formation_data:               "",
				Responsible_person_job_title_genetive: "",
				Category:                              "",
			}

			for i, value := range row {
				err = structutils.SetFieldValue(&counterparty, entity.CounterpartyRuEngMap[headers[i]], value)
				if err != nil {
					return nil, err
				}
			}

			counterparties = append(counterparties, &counterparty)
		}
	}

	return &counterparties, nil
}

func saveCounterparties(reader io.Reader) (*entity.Counterparties, error) {
	file, err := excelize.OpenReader(reader)

	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		return nil, err
	}

	counterparties := entity.Counterparties{}

	if len(rows) > 0 {
		headers := rows[0]

		for _, row := range rows[1:] {
			//TODO подумать
			counterparty := entity.Counterparty{
				Code_ou:                               "",
				Inn:                                   "",
				Institution_short_name:                "",
				Institution_full_name:                 "",
				Address:                               "",
				City:                                  "",
				Bank_details:                          "",
				Responsible_person_job_title:          "",
				Responsible_person_short_name:         "",
				Responsible_person_full_name:          "",
				Responsible_person_full_name_genitive: "",
				Acting_on:                             "",
				Ikz_2025:                              "",
				Source_funding:                        "",
				Email:                                 "",
				Phone_number:                          "",
				Contract_form:                         "",
				Contract_type:                         "",
				Contract_number:                       "",
				Contract_formation_data:               "",
				Responsible_person_job_title_genetive: "",
				Category:                              "",
			}

			for i, value := range row {
				err = structutils.SetFieldValue(&counterparty, entity.CounterpartyRuEngMap[headers[i]], value)
				if err != nil {
					return nil, err
				}
			}

			counterparties = append(counterparties, &counterparty)
		}
	}

	return &counterparties, nil
}
