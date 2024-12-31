package customui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Кастомный элемент списка

type MyListItemWidget struct {
	widget.BaseWidget
	Check   *widget.Check
	Title   *widget.Label
	Comment *widget.Label
}

func NewMyListItemWidget(title, comment string, checkChanged func(bool)) *MyListItemWidget {
	item := &MyListItemWidget{
		Check:   widget.NewCheck("", checkChanged),
		Title:   widget.NewLabel(title),
		Comment: widget.NewLabel(comment),
	}
	// item.Title.Truncation = fyne.TextTruncateEllipsis
	item.ExtendBaseWidget(item)

	return item
}

func (item *MyListItemWidget) CreateRenderer() fyne.WidgetRenderer {
	// c := container.NewBorder(nil, nil, nil, item.Check, item.Comment, item.Title)
	c := container.NewHBox(item.Check, item.Title, item.Comment)
	return widget.NewSimpleRenderer(c)
}

func (item *MyListItemWidget) SetText(titleText string, commentText string) {
	item.Title.Text = titleText
	item.Comment.Text = commentText
	item.Refresh()
}

// Лейбл с обработкой нажатия
type TappableLabel struct {
	widget.Label
}

func newTappableLabel(text string) *TappableLabel {
	label := &TappableLabel{}
	label.SetText(text)
	return label
}

// Реализация метода Tapped для интерфейса fyne.Tappable
func (l *TappableLabel) Tapped(*fyne.PointEvent) {
	fmt.Printf("Вы выбрали элемент: %s\n", l.Text)
}
