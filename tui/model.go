package tui

import (
	"time"
	"tuidoo/entities"
	"tuidoo/managers"
	"tuidoo/services"
	"tuidoo/tui/components/footer"
	"tuidoo/tui/components/menu"
	"tuidoo/tui/components/themelist"
	"tuidoo/tui/components/todoform"
	"tuidoo/tui/components/todolist"
	"tuidoo/tui/context"
	"tuidoo/tui/keys"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/charmbracelet/log"
)

type View int

const (
	ViewMain View = iota
	ViewThemes
	ViewTodoEdit
	ViewProjects
)

type Model struct {
	ctx         *context.ProgramContext
	keys        *keys.KeyMap
	currentView View

	// Components
	menu      menu.Model
	todoList  todolist.Model
	themeList themelist.Model
	todoForm  todoform.Model
	footer    footer.Model

	// State
	selectedTodo  *entities.ToDo
	focusedOnMenu bool
	taskSpinner   spinner.Model
	tasks         map[string]context.Task
}

func NewModel(sc *services.ServiceCollection, tm *managers.ThemeManager) Model {
	taskSpinner := spinner.Model{Spinner: spinner.Dot}

	ctx := &context.ProgramContext{
		Services:     sc,
		ThemeManager: tm,
		StartTask: func(task context.Task) tea.Cmd {
			log.Info("Starting task", "id", task.Id)
			task.StartTime = time.Now()
			return taskSpinner.Tick
		},
	}

	m := Model{
		ctx:           ctx,
		keys:          keys.Keys,
		currentView:   ViewMain,
		focusedOnMenu: true,
		taskSpinner:   taskSpinner,
		tasks:         map[string]context.Task{},
	}

	m.menu = menu.NewModel(ctx)
	m.todoList = todolist.NewModel(ctx)
	m.themeList = themelist.NewModel(ctx)
	m.todoForm = todoform.NewModel(ctx)
	m.footer = footer.NewModel(ctx)

	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.initScreen,
		tea.EnterAltScreen,
		m.todoList.FetchTodos(),
	)
}

func (m *Model) initScreen() tea.Msg {
	return initMsg{}
}

type initMsg struct{}
