package customui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Кастомный элемент списка c заголовком и множеством лейблов
type ListItemTitleLabels struct {
	widget.BaseWidget
	Check     *widget.Check
	Title     *widget.Label
	Labels    *widget.Label
	HandleTap func()
}

func NewListItemTitleLabels(title string, labels []string, checkChanged func(bool), handleTap func()) *ListItemTitleLabels {
	var strlabel string
	for _, label := range labels {
		strlabel += label + "          |          "
	}
	item := &ListItemTitleLabels{
		Check:     widget.NewCheck("", checkChanged),
		Title:     widget.NewLabel(title),
		Labels:    widget.NewLabel(strlabel),
		HandleTap: handleTap,
	}
	// item.Title.Truncation = fyne.TextTruncateEllipsis
	item.ExtendBaseWidget(item)

	return item
}

// Fyne package implementation
func (item *ListItemTitleLabels) CreateRenderer() fyne.WidgetRenderer {
	// c := container.NewBorder(nil, nil, nil, item.Check, item.Comment, item.Title)
	content := []fyne.CanvasObject{item.Check, item.Title, item.Labels}

	c := container.NewHBox(content...)
	return widget.NewSimpleRenderer(c)
}

func (item *ListItemTitleLabels) Tapped(*fyne.PointEvent) {
	if item.HandleTap != nil {
		item.HandleTap()
	}
}

// Component accessor's
func (item *ListItemTitleLabels) SetCheckChange(checkChanged func(bool)) {
	item.Check.OnChanged = checkChanged
}

func (item *ListItemTitleLabels) SetTapHandler(handleTap func()) {
	item.HandleTap = handleTap
}

func (litl *ListItemTitleLabels) SetLabels(labels []string) {
	var strlabel string
	for _, label := range labels {
		strlabel += label + "          |          "
	}
	litl.Labels.SetText(strlabel)
}
