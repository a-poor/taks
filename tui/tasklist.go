package tui

import (
	"github.com/a-poor/taks/lib"
	"github.com/charmbracelet/bubbles/list"
)

type listItem struct {
	task *lib.Task
}

func NewListItem(t *lib.Task) list.Item {
	return listItem{t}.toBubbleItem()
}

func (i listItem) Title() string {
	return i.task.Name
}

func (i listItem) Description() string {
	return i.task.Details
}

func (i listItem) FilterValue() string {
	return i.task.Name
}

func (i listItem) toBubbleItem() list.Item {
	return list.Item(i)
}
