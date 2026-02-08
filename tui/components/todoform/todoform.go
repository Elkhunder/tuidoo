package todoform

import (
	"fmt"
	"strings"
	"tuidoo/entities"
	"tuidoo/enums"
	"tuidoo/tui/context"
	"tuidoo/tui/keys"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx        *context.ProgramContext
	todo       *entities.ToDo
	nameInput  textinput.Model
	descInput  textarea.Model
	focusIndex int
	inputs     []string
}

type TodoSavedMsg struct {
	Todo *entities.ToDo
}

func NewModel(ctx *context.ProgramContext) Model {
	nameInput := textinput.New()
	nameInput.Placeholder = "Task name"
	nameInput.CharLimit = 100
	nameInput.Width = 50

	descInput := textarea.New()
	descInput.Placeholder = "Description..."
	descInput.SetHeight(5)
	descInput.SetWidth(50)

	return Model{
		ctx:        ctx,
		nameInput:  nameInput,
		descInput:  descInput,
		focusIndex: 0,
		inputs:     []string{"name", "description", "priority", "status", "done"},
	}
}

func (m *Model) SetTodo(todo *entities.ToDo) {
	m.todo = todo
	m.nameInput.SetValue(todo.Name)
	m.nameInput.Focus()

	desc := ""
	if todo.Description != nil {
		desc = *todo.Description
	} else if todo.Details != nil {
		desc = *todo.Details
	}
	m.descInput.SetValue(desc)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+s":
			return m, m.saveTodo()

		case "tab", "shift+tab":
			if msg.String() == "tab" {
				m.focusIndex++
			} else {
				m.focusIndex--
			}

			if m.focusIndex > len(m.inputs)-1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}

			if m.focusIndex == 0 {
				m.nameInput.Focus()
				m.descInput.Blur()
			} else if m.focusIndex == 1 {
				m.nameInput.Blur()
				m.descInput.Focus()
			} else {
				m.nameInput.Blur()
				m.descInput.Blur()
			}

			return m, nil
		}
	}

	// Update focused input
	if m.focusIndex == 0 {
		m.nameInput, cmd = m.nameInput.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.focusIndex == 1 {
		m.descInput, cmd = m.descInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.todo == nil {
		return "No todo selected"
	}

	theme := m.ctx.ThemeManager.GetCurrentTheme()

	titleStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Primary)).
		Bold(true).
		Padding(1, 1)

	labelStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Foreground)).
		Bold(true)

	valueStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.TextSecondary))

	helpStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.TextSecondary)).
		Padding(1, 0)

	buttonStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Background)).
		Background(context.TcellToLipgloss(theme.Colors.Primary)).
		Padding(0, 2).
		MarginTop(1)

	var s strings.Builder

	s.WriteString(titleStyle.Render(fmt.Sprintf("Edit Todo: %s", m.todo.Name)))
	s.WriteString("\n\n")

	// Name field
	s.WriteString(labelStyle.Render("Task Name:"))
	s.WriteString("\n")
	s.WriteString(m.nameInput.View())
	s.WriteString("\n\n")

	// Description field
	s.WriteString(labelStyle.Render("Description:"))
	s.WriteString("\n")
	s.WriteString(m.descInput.View())
	s.WriteString("\n\n")

	// Priority (read-only for now, could add selector)
	s.WriteString(labelStyle.Render("Priority: "))
	priorityColor := getPriorityColor(m.todo.Priority.String(), theme)
	priorityStyle := lipgloss.NewStyle().Foreground(priorityColor).Bold(true)
	s.WriteString(priorityStyle.Render(m.todo.Priority.String()))
	s.WriteString("\n\n")

	// Status (read-only for now)
	s.WriteString(labelStyle.Render("Status: "))
	s.WriteString(valueStyle.Render(m.todo.Status.String()))
	s.WriteString("\n\n")

	// Done status
	doneIcon := "☐"
	if m.todo.Done {
		doneIcon = "☑"
	}
	s.WriteString(labelStyle.Render("Completed: "))
	s.WriteString(valueStyle.Render(doneIcon))
	s.WriteString("\n\n")

	// Project and List info
	s.WriteString(labelStyle.Render("Project: "))
	s.WriteString(valueStyle.Render(m.todo.Project.Name))
	s.WriteString("  ")
	s.WriteString(labelStyle.Render("List: "))
	s.WriteString(valueStyle.Render(m.todo.ToDoList.Name))
	s.WriteString("\n\n")

	// Buttons
	s.WriteString(buttonStyle.Render("Save (Ctrl+S)"))
	s.WriteString("  ")
	s.WriteString(helpStyle.Render("Esc: Cancel"))
	s.WriteString("\n\n")

	s.WriteString(helpStyle.Render("Tab: Next field | Shift+Tab: Previous field"))

	return s.String()
}

func (m Model) saveTodo() tea.Cmd {
	return func() tea.Msg {
		if m.todo == nil {
			return nil
		}

		// Update todo fields
		m.todo.Name = m.nameInput.Value()
		desc := m.descInput.Value()
		m.todo.Description = &desc

		// Save to database
		err := m.ctx.Services.ToDoService.Update(m.todo)
		if err != nil {
			// Could return error message
			return nil
		}

		return TodoSavedMsg{Todo: m.todo}
	}
}

func (m *Model) ApplyTheme() {
	// Theme applied on next render
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}

func getPriorityColor(priority string, theme *entities.Theme) lipgloss.Color {
	switch priority {
	case "High", "Urgent":
		return context.TcellToLipgloss(theme.Colors.Error)
	case "Medium":
		return context.TcellToLipgloss(theme.Colors.Warning)
	case "Low":
		return context.TcellToLipgloss(theme.Colors.Info)
	default:
		return context.TcellToLipgloss(theme.Colors.TextSecondary)
	}
}
