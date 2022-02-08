package tui

import (
	"github.com/a-poor/taks/lib"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	db *lib.TaksDB // Stores the connection to the task DB
}

// NewModel creates a new bubbletea TUI model.
func NewModel(db *lib.TaksDB) *model {
	return &model{
		db: db,
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return "Hello, World!"
}
