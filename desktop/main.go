package main

import (
	entity "desktop-templater-docx/internal/domain"
	"desktop-templater-docx/pkg/customui"
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

	// list := widget.NewList(
	// 	func() int { return len(counterpartiesList) },
	// 	func() fyne.CanvasObject {
	// 		// return widget.NewLabel("test")
	// 		// return widget.NewCheck("test", func(changed bool) {})
	// 		// return newTappableLabel("test")
	// 		return customui.NewMyListItemWidget("test title", "test comment", func(changed bool) {})
	// 	},
	// 	func(i int, o fyne.CanvasObject) {
	// 		comment := counterpartiesList[i].Institution_short_name + " | " + counterpartiesList[i].Responsible_person_short_name + " | " + counterpartiesList[i].City
	// 		o.(*customui.MyListItemWidget).SetText(counterpartiesList[i].Inn, comment)
	// 	},
	// )

	listContainer := container.NewVBox()

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
	button := widget.NewButton("Загрузить файл", func() {
		fileDialog := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, err error) {
				if err != nil {
					println("Ошибка при открытии файла:", err.Error())
					return
				}
				if uc != nil {
					counterparties, err := loadCounterparties(uc)
					if err != nil {
						fmt.Printf("Programm stop with error [%v]", err)
					}

					println("Выбран файл:", uc.URI().Path())
					counterpartiesList = append(counterpartiesList, *counterparties...)
					for _, c := range counterpartiesList {
						comment := c.Institution_short_name + " | " + c.Responsible_person_short_name + " | " + c.City
						listContainer.Add(customui.NewMyListItemWidget(c.Inn, comment, func(changed bool) {}))
					}
					for _, c := range counterpartiesList {
						comment := c.Institution_short_name + " | " + c.Responsible_person_short_name + " | " + c.City
						listContainer.Add(customui.NewMyListItemWidget(c.Inn, comment, func(changed bool) {}))
					}
					for _, c := range counterpartiesList {
						comment := c.Institution_short_name + " | " + c.Responsible_person_short_name + " | " + c.City
						listContainer.Add(customui.NewMyListItemWidget(c.Inn, comment, func(changed bool) {}))
					}

					listContainer.Refresh()
					// list.Refresh()
					uc.Close()
				} else {
					println("Файл не выбран")
				}
			}, w)

		fileDialog.Show()
	})

	processTemplateButton := widget.NewButton("Сформировать договора", func() {})

	// content := container.NewBorder(container.NewVBox(button), nil, nil, nil, container.NewScroll(list))
	content := container.NewBorder(container.NewVBox(button), container.NewVBox(processTemplateButton), nil, nil, container.NewScroll(listContainer))

	w.SetContent(content)

	w.Resize(fyne.NewSize(1280, 720))
	w.ShowAndRun()
}

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Custom List with Checkboxes and Buttons")

// 	// Исходные данные для списка
// 	data := []string{
// 		"Item 1", "Item 2", "Item 3", "Item 4", "Item 5",
// 	}

// 	// Список для хранения состояний чекбоксов
// 	checkboxes := make([]*widget.Check, len(data))
// 	buttons := make([]*widget.Button, len(data))

// 	// Функция для обработки клика по кнопке
// 	handleClick := func(item string) {
// 		fmt.Printf("Button clicked for: %s\n", item)
// 	}

// 	// Создание контейнера с элементами
// 	var listItems []fyne.CanvasObject
// 	for i, item := range data {
// 		// Создаем чекбокс
// 		checkboxes[i] = widget.NewCheck(item, nil)

// 		// Создаем кнопку
// 		buttons[i] = widget.NewButton("Click Me", func() {
// 			handleClick(item) // Обрабатываем клик по кнопке
// 		})

// 		// Контейнер для чекбокса и кнопки
// 		listItems = append(listItems, container.NewHBox(checkboxes[i], buttons[i]))
// 	}

// 	// Размещение элементов в контейнере VBox
// 	listContainer := container.NewVBox(listItems...)

// 	// Устанавливаем содержимое окна
// 	myWindow.SetContent(listContainer)

// 	// Показываем окно
// 	myWindow.ShowAndRun()
// }

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
