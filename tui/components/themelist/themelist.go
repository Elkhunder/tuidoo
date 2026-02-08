package themelist

import (
	"strings"
	"tuidoo/tui/context"
	"tuidoo/tui/keys"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx        *context.ProgramContext
	list       list.Model
	themeNames []string
	cursor     int
}

type ThemeChangedMsg struct {
	ThemeName string
}

type themeItem struct {
	name      string
	isCurrent bool
}

func (i themeItem) FilterValue() string { return i.name }
func (i themeItem) Title() string {
	if i.isCurrent {
		return "✓ " + i.name
	}
	return "  " + i.name
}
func (i themeItem) Description() string { return "" }

func NewModel(ctx *context.ProgramContext) Model {
	themeNames := ctx.ThemeManager.GetThemeNames()
	currentTheme := ctx.ThemeManager.GetCurrentTheme().Name

	items := make([]list.Item, 0, len(themeNames))
	for _, name := range themeNames {
		items = append(items, themeItem{
			name:      name,
			isCurrent: name == currentTheme,
		})
	}

	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Select Theme"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return Model{
		ctx:        ctx,
		list:       l,
		themeNames: themeNames,
		cursor:     0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width-40, msg.Height-10)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, keys.Keys.Down):
			if m.cursor < len(m.themeNames)-1 {
				m.cursor++
			}

		case key.Matches(msg, keys.Keys.Enter):
			selectedTheme := m.themeNames[m.cursor]
			return m, func() tea.Msg {
				return ThemeChangedMsg{ThemeName: selectedTheme}
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	theme := m.ctx.ThemeManager.GetCurrentTheme()
	currentTheme := theme.Name

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

	helpStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.TextSecondary)).
		Padding(1, 1)

	var s strings.Builder
	s.WriteString(titleStyle.Render("Theme Selector"))
	s.WriteString("\n\n")

	for i, themeName := range m.themeNames {
		cursor := "  "
		checkmark := "  "

		if i == m.cursor {
			cursor = "› "
		}

		if themeName == currentTheme {
			checkmark = "✓ "
		}

		line := checkmark + themeName

		if i == m.cursor {
			s.WriteString(cursor + selectedStyle.Render(line))
		} else {
			s.WriteString(cursor + normalStyle.Render(line))
		}
		s.WriteString("\n")
	}

	s.WriteString("\n")
	s.WriteString(helpStyle.Render("enter: apply theme | esc: back | t: toggle themes"))

	return s.String()
}

func (m *Model) ApplyTheme() {
	currentTheme := m.ctx.ThemeManager.GetCurrentTheme().Name
	items := make([]list.Item, 0, len(m.themeNames))
	for _, name := range m.themeNames {
		items = append(items, themeItem{
			name:      name,
			isCurrent: name == currentTheme,
		})
	}
	m.list.SetItems(items)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
