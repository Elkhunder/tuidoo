package todolist

import (
	"strings"
	"tuidoo/entities"
	"tuidoo/tui/context"
	"tuidoo/tui/keys"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx    *context.ProgramContext
	table  table.Model
	todos  []entities.ToDo
	width  int
	height int
}

type TodosLoadedMsg struct {
	Todos []entities.ToDo
}

type TodoSelectedMsg struct {
	Todo *entities.ToDo
}

type TodoToggledMsg struct {
	TodoId uint
}

func NewModel(ctx *context.ProgramContext) Model {
	columns := []table.Column{
		{Title: "✓", Width: 3},
		{Title: "Priority", Width: 10},
		{Title: "Task", Width: 35},
		{Title: "Project", Width: 15},
		{Title: "List", Width: 15},
		{Title: "Status", Width: 12},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	return Model{
		ctx:    ctx,
		table:  t,
		todos:  []entities.ToDo{},
		height: 20,
	}
}

func (m Model) Init() tea.Cmd {
	return m.FetchTodos()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case TodosLoadedMsg:
		m.todos = msg.Todos
		m.updateTableRows()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Enter):
			if len(m.todos) > 0 {
				idx := m.table.Cursor()
				if idx < len(m.todos) {
					return m, func() tea.Msg {
						return TodoSelectedMsg{Todo: &m.todos[idx]}
					}
				}
			}

		case key.Matches(msg, keys.Keys.Space):
			if len(m.todos) > 0 {
				idx := m.table.Cursor()
				if idx < len(m.todos) {
					return m, m.toggleTodo(&m.todos[idx])
				}
			}
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	theme := m.ctx.ThemeManager.GetCurrentTheme()

	if len(m.todos) == 0 {
		emptyStyle := lipgloss.NewStyle().
			Foreground(context.TcellToLipgloss(theme.Colors.TextSecondary)).
			Padding(2, 0).
			Align(lipgloss.Center)

		return emptyStyle.Render("No todos yet - Press 'n' to create a new task")
	}

	m.applyTableTheme()

	titleStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.Primary)).
		Bold(true).
		Padding(1, 1)

	helpStyle := lipgloss.NewStyle().
		Foreground(context.TcellToLipgloss(theme.Colors.TextSecondary)).
		Padding(1, 1)

	var s strings.Builder
	s.WriteString(titleStyle.Render("TUIDOO - Todo List"))
	s.WriteString("\n")
	s.WriteString(m.table.View())
	s.WriteString("\n")
	s.WriteString(helpStyle.Render("enter: edit | space: toggle done | n: new | r: refresh"))

	return s.String()
}

func (m *Model) updateTableRows() {
	rows := make([]table.Row, 0, len(m.todos))

	for _, todo := range m.todos {
		rows = append(rows, table.Row{
			getStatusIcon(todo.Done),
			getPriorityIcon(todo.Priority.String()) + " " + todo.Priority.String(),
			todo.Name,
			todo.Project.Name,
			todo.ToDoList.Name,
			todo.Status.String(),
		})
	}

	m.table.SetRows(rows)
}

func (m *Model) applyTableTheme() {
	theme := m.ctx.ThemeManager.GetCurrentTheme()

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(context.TcellToLipgloss(theme.Colors.Border)).
		BorderBottom(true).
		Bold(true).
		Foreground(context.TcellToLipgloss(theme.Colors.Primary))

	s.Selected = s.Selected.
		Foreground(context.TcellToLipgloss(theme.Colors.Background)).
		Background(context.TcellToLipgloss(theme.Colors.Primary)).
		Bold(false)

	s.Cell = s.Cell.
		Foreground(context.TcellToLipgloss(theme.Colors.Foreground))

	m.table.SetStyles(s)
}

func (m Model) FetchTodos() tea.Cmd {
	return func() tea.Msg {
		todos, err := m.ctx.Services.ToDoService.GetAll(true)
		if err != nil {
			return TodosLoadedMsg{Todos: []entities.ToDo{}}
		}
		return TodosLoadedMsg{Todos: todos}
	}
}

func (m Model) toggleTodo(todo *entities.ToDo) tea.Cmd {
	return func() tea.Msg {
		var err error
		if !todo.Done {
			err = m.ctx.Services.ToDoService.MarkAsComplete(todo.ID)
		} else {
			err = m.ctx.Services.ToDoService.MarkAsIncomplete(todo.ID)
		}

		if err != nil {
			return TodosLoadedMsg{Todos: m.todos}
		}

		// Reload todos
		todos, err := m.ctx.Services.ToDoService.GetAll(true)
		if err != nil {
			return TodosLoadedMsg{Todos: m.todos}
		}

		return TodosLoadedMsg{Todos: todos}
	}
}

func (m *Model) ApplyTheme() {
	m.applyTableTheme()
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}

func getStatusIcon(done bool) string {
	if done {
		return "✓"
	}
	return " "
}

func getPriorityIcon(priority string) string {
	switch priority {
	case "High", "Urgent":
		return "🔥"
	case "Medium":
		return "⚡"
	case "Low":
		return "💤"
	default:
		return "  "
	}
}
