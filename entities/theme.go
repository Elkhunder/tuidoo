package entities

import "github.com/gdamore/tcell/v2"

type Colors struct {
	// Header gradient (5 colors for ASCII art gradient)
	HeaderGradient [6]tcell.Color

	// Primary colors
	Primary      tcell.Color
	PrimaryDark  tcell.Color
	PrimaryLight tcell.Color

	// Accent colors
	Accent      tcell.Color
	AccentDark  tcell.Color
	AccentLight tcell.Color

	// Status colors
	Success tcell.Color
	Warning tcell.Color
	Error   tcell.Color
	Info    tcell.Color

	// Neutral colors
	Background tcell.Color
	Foreground tcell.Color // Added: Main foreground/text color
	Surface    tcell.Color
	Border     tcell.Color

	// Text colors
	TextPrimary   tcell.Color
	TextSecondary tcell.Color
	TextDisabled  tcell.Color
	TextInverse   tcell.Color

	// Interactive states
	Hover    tcell.Color
	Active   tcell.Color
	Focus    tcell.Color
	Selected tcell.Color
}

type Theme struct {
	ID     string
	Name   string
	Colors Colors
}

func (t *Theme) PriorityColor(priority string) tcell.Color {
	switch priority {
	case "high", "critical":
		return t.Colors.Error
	case "medium":
		return t.Colors.Warning
	case "low":
		return t.Colors.Info
	default:
		return t.Colors.TextSecondary
	}
}

func (t *Theme) StatusColor(status string) tcell.Color {
	switch status {
	case "done", "complete", "completed":
		return t.Colors.Success
	case "in-progress", "active":
		return t.Colors.Info
	case "blocked", "failed":
		return t.Colors.Error
	case "pending":
		return t.Colors.Warning
	default:
		return t.Colors.TextPrimary
	}
}
