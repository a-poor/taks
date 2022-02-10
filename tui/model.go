package tui

import (
	"github.com/a-poor/taks/lib"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	db   *lib.TaksDB // Stores the connection to the task DB
	list list.Model
}

// NewModel creates a new bubbletea TUI model.
func NewModel(db *lib.TaksDB) (*model, error) {
	// Get the tasks
	tasks, err := db.ListTasks()
	if err != nil {
		return nil, err
	}

	// Convert the tasks to list items
	items := make([]list.Item, len(tasks))
	for i, task := range tasks {
		items[i] = NewListItem(task)
	}

	// Using the default delegate
	del := list.NewDefaultDelegate()

	// Initial width and height
	w, h := 0, 0

	// Create the list model
	l := list.New(items, del, w, h)

	// Configure the list model
	l.Title = "ToDo Tasks"

	return &model{
		db:   db,
		list: l,
	}, nil
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
