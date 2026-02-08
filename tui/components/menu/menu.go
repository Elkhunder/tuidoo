package menu

import (
	"strings"
	"tuidoo/tui/context"
	"tuidoo/tui/keys"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuAction struct {
	View string
}

type MenuItem struct {
	Label       string
	Description string
	Key         rune
	View        string
}

type Model struct {
	ctx            *context.ProgramContext
	items          []MenuItem
	cursor         int
	selectedAction *MenuAction
}

func NewModel(ctx *context.ProgramContext) Model {
	items := []MenuItem{
		{Label: "View ToDos", Description: "View all todos", Key: 'l', View: "main"},
		{Label: "New Task", Description: "Create new task", Key: 'n', View: "new"},
		{Label: "Projects", Description: "Manage projects", Key: 'p', View: "projects"},
		{Label: "Settings", Description: "App settings", Key: 't', View: "themes"},
		{Label: "Quit", Description: "Exit application", Key: 'q', View: "quit"},
	}

	return Model{
		ctx:    ctx,
		items:  items,
		cursor: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, keys.Keys.Down):
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case key.Matches(msg, keys.Keys.Enter):
			selected := m.items[m.cursor]
			if selected.View == "quit" {
				return m, tea.Quit
			}
			m.selectedAction = &MenuAction{View: selected.View}
		}
	}

	return m, nil
}

func (m Model) View() string {
	theme := m.ctx.ThemeManager.GetCurrentTheme()

	selectedStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Background)).
		Background(context.TcellToLipgloss(theme.Colors.Primary)).
		Bold(true).
		Padding(0, 1)

	normalStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Foreground)).
		Padding(0, 1)

	titleStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Primary)).
		Bold(true).
		Padding(1, 1)

	var s strings.Builder
	s.WriteString(titleStyle.Render("Menu"))
	s.WriteString("\n\n")

	for i, item := range m.items {
		cursor := "  "
		if i == m.cursor {
			cursor = "› "
		}

		if i == m.cursor {
			s.WriteString(cursor + selectedStyle.Render(item.Label))
		} else {
			s.WriteString(cursor + normalStyle.Render(item.Label))
		}
		s.WriteString("\n")
	}

	return s.String()
}

func (m Model) GetSelectedAction() *MenuAction {
	return m.selectedAction
}

func (m *Model) ClearAction() {
	m.selectedAction = nil
}

func (m *Model) ApplyTheme() {
	// Theme will be applied on next render
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
