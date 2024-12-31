package customui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Кастомный элемент списка

type MyListItemWidget struct {
	widget.BaseWidget
	Check     *widget.Check
	Title     *widget.Label
	Comment   *widget.Label
	HandleTap func()
}

func NewMyListItemWidget(title, comment string, checkChanged func(bool), handleTap func()) *MyListItemWidget {
	item := &MyListItemWidget{
		Check:     widget.NewCheck("", checkChanged),
		Title:     widget.NewLabel(title),
		Comment:   widget.NewLabel(comment),
		HandleTap: handleTap,
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

func (item *MyListItemWidget) Tapped(*fyne.PointEvent) {
	if item.HandleTap != nil {
		item.HandleTap()
	}
}
