package context

import (
	"fmt"
	"time"
	"tuidoo/managers"
	"tuidoo/services"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gdamore/tcell/v2"
)

type ProgramContext struct {
	Services     *services.ServiceCollection
	ThemeManager *managers.ThemeManager

	ScreenWidth       int
	ScreenHeight      int
	MainContentWidth  int
	MainContentHeight int

	StartTask func(task Task) tea.Cmd
}

type TaskState int

const (
	TaskStart TaskState = iota
	TaskFinished
	TaskError
)

type Task struct {
	Id           string
	StartTime    time.Time
	FinishedTime *time.Time
	State        TaskState
	StartText    string
	FinishedText string
	Error        error
}

// Helper to convert tcell.Color to lipgloss.Color
func TcellToLipgloss(c tcell.Color) lipgloss.Color {
	r, g, b := c.RGB()
	return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", r, g, b))
}
