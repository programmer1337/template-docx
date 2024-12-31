package customui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Лейбл с обработкой нажатия
type TappableLabel struct {
	widget.Label
}

func newTappableLabel(text string) *TappableLabel {
	label := &TappableLabel{}
	label.SetText(text)
	return label
}

func (l *TappableLabel) Tapped(*fyne.PointEvent) {
	fmt.Printf("Вы выбрали элемент: %s\n", l.Text)
}
