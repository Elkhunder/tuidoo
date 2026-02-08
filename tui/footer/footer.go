package footer

import (
	"strings"
	"tuidoo/tui/context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx   *context.ProgramContext
	width int
}

func NewModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:   ctx,
		width: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}

func (m Model) View() string {
	theme := m.ctx.ThemeManager.GetCurrentTheme()

	helpStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.TextSecondary)).
		Background(context.TcellToLipgloss(theme.Colors.Surface)).
		Padding(0, 1)

	keyStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Primary)).
		Bold(true)

	var helpItems []string
	helpItems = append(helpItems,
		keyStyle.Render("↑/k")+" up",
		keyStyle.Render("↓/j")+" down",
		keyStyle.Render("enter")+" select",
		keyStyle.Render("space")+" toggle done",
		keyStyle.Render("tab")+" switch focus",
		keyStyle.Render("t")+" themes",
		keyStyle.Render("r")+" refresh",
		keyStyle.Render("n")+" new task",
		keyStyle.Render("q")+" quit",
	)

	help := strings.Join(helpItems, " • ")

	if m.width > 0 {
		return helpStyle.Width(m.width).Render(help)
	}

	return helpStyle.Render(help)
}

func (m *Model) ApplyTheme() {
	// Theme applied on next render
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}

func (m *Model) SetWidth(width int) {
	m.width = width
}
