package customui

import (
	"desktop-templater-docx/internal/domain/entity"
	"desktop-templater-docx/pkg/sliceutils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Кастомный элемент списка

type ListTitleLabels struct {
	widget.BaseWidget
	ListItems      []*ListItemTitleLabels
	SelectedItems  []int
	Container      *fyne.Container
	selectAllCheck *widget.Check
	createItem     func(pos int, item *entity.Counterparty) *ListItemTitleLabels
}

func NewListTitleLabels(handleTap func(pos int, item *ListItemTitleLabels), handleChange func(isSelected bool)) *ListTitleLabels {
	item := &ListTitleLabels{
		ListItems:     []*ListItemTitleLabels{},
		SelectedItems: []int{},
		Container:     container.NewVBox(),
	}
	item.selectAllCheck = item.CreateCheckAll()
	item.createItem = item.CreateItem(handleTap, handleChange)
	item.ExtendBaseWidget(item)

	return item
}

func (ltl *ListTitleLabels) CreateRenderer() fyne.WidgetRenderer {
	// c := container.NewBorder(nil, nil, nil, item.Check, item.Comment, item.Title)
	for _, item := range ltl.ListItems {
		ltl.Container.Add(item)
	}

	return widget.NewSimpleRenderer(ltl.Container)
}

func (ltl *ListTitleLabels) CreateCheckAll() *widget.Check {
	checkAll := widget.NewCheck("Выбрать все", func(changed bool) {
		if changed {
			ltl.SelectAll()
		} else {
			if len(ltl.SelectedItems) == len(ltl.ListItems) {
				ltl.UnselectAll()
			}
		}
	})

	ltl.Container.Add(checkAll)
	checkAll.Hide()

	return checkAll
}

func (ltl *ListTitleLabels) CreateItem(handleClick func(pos int, item *ListItemTitleLabels), handleChange func(isSelected bool)) func(pos int, c *entity.Counterparty) *ListItemTitleLabels {
	return func(pos int, c *entity.Counterparty) *ListItemTitleLabels {
		return NewListItemTitleLabels(
			c.Inn,
			[]string{c.Institution_short_name, c.Responsible_person_short_name, c.City},
			func(isSelected bool) {
				ltl.handleSelect(pos, isSelected)
				if handleChange != nil {
					handleChange(isSelected)
				}
			},
			func() {
				fmt.Printf("Was tapped: %v\n", pos)
				if handleClick != nil {
					handleClick(pos, ltl.ListItems[pos])
				}
			})
	}
}

func (ltl *ListTitleLabels) SetItems(counterparties entity.Counterparties) {
	ltl.Container.RemoveAll()
	ltl.ListItems = []*ListItemTitleLabels{}

	for pos, c := range counterparties {
		ltl.ListItems = append(ltl.ListItems, ltl.createItem(pos, c))
	}

	ltl.Container.Add(ltl.selectAllCheck)
	for _, listItem := range ltl.ListItems {
		ltl.Container.Add(listItem)
	}

	ltl.controlSelectAllVisible()
}

func (ltl *ListTitleLabels) AddItems(counterparties *entity.Counterparties) {
	for pos, c := range *counterparties {
		ltl.ListItems = append(ltl.ListItems, ltl.createItem(pos, c))
	}

	for _, listItem := range ltl.ListItems {
		ltl.Container.Add(listItem)
	}
}

func (ltl *ListTitleLabels) GetSelected() []int {
	return ltl.SelectedItems
}

func (ltl *ListTitleLabels) SelectAll() {
	for i := range len(ltl.ListItems) {
		ltl.ListItems[i].Check.SetChecked(true)
	}

	fmt.Printf("Selected %v\n", ltl.SelectedItems)
}

func (ltl *ListTitleLabels) UnselectAll() {
	for i := range len(ltl.ListItems) {
		ltl.ListItems[i].Check.SetChecked(false)
	}

	fmt.Printf("Selected %v\n", ltl.SelectedItems)
}

func (ltl *ListTitleLabels) handleSelect(id int, isSelected bool) {
	if isSelected {
		ltl.SelectedItems = append(ltl.SelectedItems, id)
	} else {
		ltl.SelectedItems = sliceutils.RemoveByValue[int](ltl.SelectedItems, id)
		ltl.selectAllCheck.SetChecked(false)
	}
}

func (ltl *ListTitleLabels) controlSelectAllVisible() {
	if len(ltl.ListItems) > 0 {
		ltl.selectAllCheck.Show()
	}
}
