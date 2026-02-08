package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"tuidoo/tui/context"
)

func (m Model) View() string {
	if m.ctx.ThemeManager == nil {
		return "Initializing..."
	}

	theme := m.ctx.ThemeManager.GetCurrentTheme()

	// Header
	headerStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Primary)).
		Bold(true).
		Align(lipgloss.Center).
		Padding(1, 0)

	header := headerStyle.Render(m.ctx.ThemeManager.CreateHeader())

	// Menu styling
	menuStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(context.TcellToLipgloss(theme.Colors.Border)).
		Width(28).
		Height(m.ctx.MainContentHeight)

	// Content styling
	contentStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(context.TcellToLipgloss(theme.Colors.Border)).
		Width(m.ctx.MainContentWidth).
		Height(m.ctx.MainContentHeight)

	// Render content based on view
	var content string
	switch m.currentView {
	case ViewMain:
		content = m.todoList.View()

	case ViewThemes:
		content = m.themeList.View()

	case ViewTodoEdit:
		content = m.todoForm.View()

	case ViewProjects:
		content = "Projects view - Coming soon!"
	}

	// Highlight focused component
	menuBorder := lipgloss.RoundedBorder()
	contentBorder := lipgloss.RoundedBorder()

	if m.focusedOnMenu {
		menuStyle = menuStyle.BorderForeground(context.TcellToLipgloss(theme.Colors.Primary))
	} else {
		contentStyle = contentStyle.BorderForeground(context.TcellToLipgloss(theme.Colors.Primary))
	}

	// Layout
	mainLayout := lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuStyle.Render(m.menu.View()),
		contentStyle.Render(content),
	)

	var b strings.Builder
	b.WriteString(header)
	b.WriteString("\n")
	b.WriteString(mainLayout)
	b.WriteString("\n")
	b.WriteString(m.footer.View())

	return b.String()
}
