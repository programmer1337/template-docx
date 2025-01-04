package creator

import (
	"desktop-templater-docx/internal/domain/entity"
	"desktop-templater-docx/pkg/structutils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type InputPlaceholder struct {
	Input       *widget.Entry
	Name        string
	placeholder string
}

// TODO Прокинуть поля сверху
func CreateForm(onSave func(inputs []InputPlaceholder), onClose func(), counterparty *entity.Counterparty) *fyne.Container {
	inputs := []InputPlaceholder{}

	form := widget.NewForm()
	fields := structutils.GetStructFieldNames(entity.Counterparty{})
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
		inputs = append(inputs, InputPlaceholder{Input: input, Name: fieldName, placeholder: entity.CounterpartyEngRuAlias[fieldName]})

		form.Append(entity.CounterpartyEngRuAlias[fieldName], input)
	}

	saveButton := widget.NewButton("Сохранить", func() {
		if onSave != nil {
			onSave(inputs)
		}
	})

	cancelButton := widget.NewButton("Отменить", func() {
		if onClose != nil {
			onClose()
		}
	})

	buttonContainer := container.NewHBox(saveButton, cancelButton)

	formContent := container.NewBorder(
		widget.NewLabel(fmt.Sprintf("Редактирование контрагента %v", counterparty.Inn)),
		buttonContainer,
		nil,
		nil,
		container.NewScroll(form),
	)

	return formContent
}
