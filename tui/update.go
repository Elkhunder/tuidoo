package tui

import (
	"time"
	"tuidoo/tui/components/themelist"
	"tuidoo/tui/components/todoform"
	"tuidoo/tui/components/todolist"
	"tuidoo/tui/context"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/charmbracelet/log"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		log.Info("Key pressed", "key", msg.String())

		// Global quit
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}

		// Handle escape
		if key.Matches(msg, m.keys.Escape) {
			if m.currentView == ViewTodoEdit {
				m.currentView = ViewMain
				m.focusedOnMenu = false
				return m, nil
			} else if m.currentView == ViewThemes {
				m.currentView = ViewMain
				m.focusedOnMenu = false
				return m, nil
			}
		}

		// View switching
		switch {
		case key.Matches(msg, m.keys.ToggleThemes):
			if m.currentView == ViewThemes {
				m.currentView = ViewMain
			} else {
				m.currentView = ViewThemes
			}
			m.focusedOnMenu = false
			return m, nil

		case key.Matches(msg, m.keys.Tab):
			m.focusedOnMenu = !m.focusedOnMenu
			return m, nil

		case key.Matches(msg, m.keys.Refresh):
			cmd = m.todoList.FetchTodos()
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)

	case initMsg:
		cmds = append(cmds, m.todoList.FetchTodos())

	case todolist.TodosLoadedMsg:
		m.todoList, cmd = m.todoList.Update(msg)
		cmds = append(cmds, cmd)

	case todolist.TodoSelectedMsg:
		m.selectedTodo = msg.Todo
		m.currentView = ViewTodoEdit
		m.todoForm.SetTodo(msg.Todo)
		m.focusedOnMenu = false
		return m, nil

	case themelist.ThemeChangedMsg:
		m.applyTheme(msg.ThemeName)
		return m, nil

	case todoform.TodoSavedMsg:
		m.currentView = ViewMain
		m.focusedOnMenu = false
		return m, m.todoList.FetchTodos()

	case TaskFinishedMsg:
		m.handleTaskFinished(msg)
	}

	// Update focused component
	if m.focusedOnMenu {
		m.menu, cmd = m.menu.Update(msg)
		cmds = append(cmds, cmd)

		// Check if menu action triggered view change
		if action := m.menu.GetSelectedAction(); action != nil {
			switch action.View {
			case "main":
				m.currentView = ViewMain
				m.focusedOnMenu = false
			case "themes":
				m.currentView = ViewThemes
				m.focusedOnMenu = false
			case "projects":
				m.currentView = ViewProjects
				m.focusedOnMenu = false
			}
			m.menu.ClearAction()
		}
	} else {
		switch m.currentView {
		case ViewMain:
			m.todoList, cmd = m.todoList.Update(msg)
			cmds = append(cmds, cmd)

		case ViewThemes:
			m.themeList, cmd = m.themeList.Update(msg)
			cmds = append(cmds, cmd)

		case ViewTodoEdit:
			m.todoForm, cmd = m.todoForm.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	m.footer, cmd = m.footer.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	m.ctx.MainContentHeight = msg.Height - 15 // Header + footer
	m.ctx.MainContentWidth = msg.Width - 35   // Menu width
}

func (m *Model) applyTheme(themeName string) {
	m.ctx.ThemeManager.SetTheme(themeName)
	m.ctx.Services.SettingsService.SetActiveTheme(themeName)

	// Refresh all components
	m.todoList.ApplyTheme()
	m.themeList.ApplyTheme()
	m.menu.ApplyTheme()
	m.footer.ApplyTheme()
}

func (m *Model) handleTaskFinished(msg TaskFinishedMsg) {
	task, ok := m.tasks[msg.TaskId]
	if ok {
		log.Info("Task finished", "id", task.Id)
		if msg.Err != nil {
			log.Error("Task error", "id", task.Id, "err", msg.Err)
			task.State = context.TaskError
			task.Error = msg.Err
		} else {
			task.State = context.TaskFinished
		}
		now := time.Now()
		task.FinishedTime = &now
		m.tasks[msg.TaskId] = task
	}
}

type TaskFinishedMsg struct {
	TaskId string
	Err    error
}
