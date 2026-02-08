package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"tuidoo/entities"
	"tuidoo/enums"
	"tuidoo/managers"
	"tuidoo/services"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Page constants
type page int

const (
	PageMain page = iota
	PageThemes
	PageWelcome
	PageTodoDetails
)

// ============================================================================
// Messages (Bubbletea's event system)
// ============================================================================

type todosLoadedMsg struct {
	todos []entities.ToDo
}

type todoUpdatedMsg struct {
	todo *entities.ToDo
}

type themeChangedMsg struct {
	themeName string
}

type errorMsg struct {
	err error
}

type switchPageMsg struct {
	page page
}

type toggleTodoMsg struct {
	todoIndex int
}

// ============================================================================
// Model (Application State)
// ============================================================================

type Model struct {
	services     *services.ServiceCollection
	themeManager *managers.ThemeManager

	// Current state
	currentPage  page
	todos        []entities.ToDo
	selectedTodo *entities.ToDo
	err          error

	// UI Components
	todoTable  table.Model
	themeList  list.Model
	menuCursor int
	menuItems  []menuItem

	// Todo details form fields
	nameInput        textinput.Model
	descriptionInput textarea.Model
	priorityIndex    int
	statusIndex      int
	doneCheckbox     bool

	// Window dimensions
	width  int
	height int

	// Focus tracking
	focusedComponent focusedComponent
}

type focusedComponent int

const (
	focusMenu focusedComponent = iota
	focusContent
	focusTodoForm
)

type menuItem struct {
	label       string
	description string
	action      func(*Model) tea.Cmd
}

// ============================================================================
// Initialization
// ============================================================================

func RunTUI() error {
	sc, err := services.NewServiceCollection()
	if err != nil {
		return fmt.Errorf("failed to initialize services: %w", err)
	}

	themeManager := managers.NewThemeManager()
	if err := applyUserTheme(sc, themeManager); err != nil {
		sc.Close()
		return fmt.Errorf("theme initialization failed: %w", err)
	}

	m := initialModel(sc, themeManager)

	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	if _, err := p.Run(); err != nil {
		sc.Close()
		return fmt.Errorf("error running tuidoo: %w", err)
	}

	sc.Close()
	return nil
}

func initialModel(sc *services.ServiceCollection, tm *managers.ThemeManager) *Model {
	// Initialize menu items
	menuItems := []menuItem{
		{label: "View ToDos", description: "View all todos", action: func(m *Model) tea.Cmd {
			m.currentPage = PageMain
			m.focusedComponent = focusContent
			return nil
		}},
		{label: "New Task", description: "Create new task", action: func(m *Model) tea.Cmd {
			return showNotImplemented("New Task")
		}},
		{label: "Projects", description: "Manage projects", action: func(m *Model) tea.Cmd {
			return showNotImplemented("Projects")
		}},
		{label: "Settings", description: "App settings", action: func(m *Model) tea.Cmd {
			m.currentPage = PageThemes
			m.focusedComponent = focusContent
			return nil
		}},
		{label: "Quit", description: "Exit application", action: func(m *Model) tea.Cmd {
			return tea.Quit
		}},
	}

	// Initialize todo table
	columns := []table.Column{
		{Title: "✓", Width: 3},
		{Title: "Priority", Width: 10},
		{Title: "Task", Width: 30},
		{Title: "Project", Width: 15},
		{Title: "List", Width: 15},
		{Title: "Status", Width: 12},
	}

	todoTable := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	// Initialize theme list
	themeItems := make([]list.Item, 0)
	for _, themeName := range tm.GetThemeNames() {
		themeItems = append(themeItems, themeListItem{
			name:      themeName,
			isCurrent: themeName == tm.GetCurrentTheme().Name,
		})
	}

	themeList := list.New(themeItems, themeListDelegate{}, 40, 20)
	themeList.Title = "Select Theme"
	themeList.SetShowStatusBar(false)
	themeList.SetFilteringEnabled(false)

	// Initialize form inputs
	nameInput := textinput.New()
	nameInput.Placeholder = "Task name"
	nameInput.CharLimit = 50

	descriptionInput := textarea.New()
	descriptionInput.Placeholder = "Description..."
	descriptionInput.SetHeight(5)
	descriptionInput.SetWidth(50)

	m := &Model{
		services:         sc,
		themeManager:     tm,
		currentPage:      PageMain,
		menuItems:        menuItems,
		todoTable:        todoTable,
		themeList:        themeList,
		nameInput:        nameInput,
		descriptionInput: descriptionInput,
		focusedComponent: focusMenu,
	}

	return m
}

// Init is called once when the program starts
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		loadTodos(m.services),
		textinput.Blink,
	)
}

// ============================================================================
// Update (Event Handler)
// ============================================================================

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.todoTable.SetWidth(msg.Width - 35)
		m.todoTable.SetHeight(msg.Height - 15)
		return m, nil

	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case todosLoadedMsg:
		m.todos = msg.todos
		m.updateTableRows()
		return m, nil

	case todoUpdatedMsg:
		// Reload todos after update
		return m, loadTodos(m.services)

	case themeChangedMsg:
		m.applyThemeStyles()
		return m, nil

	case errorMsg:
		m.err = msg.err
		log.Printf("Error: %v", msg.err)
		return m, nil

	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			return m.handleMouseClick(msg)
		}
	}

	// Update focused component
	var cmd tea.Cmd
	switch m.currentPage {
	case PageMain:
		if m.focusedComponent == focusContent {
			m.todoTable, cmd = m.todoTable.Update(msg)
			cmds = append(cmds, cmd)
		}
	case PageThemes:
		if m.focusedComponent == focusContent {
			m.themeList, cmd = m.themeList.Update(msg)
			cmds = append(cmds, cmd)
		}
	case PageTodoDetails:
		cmds = append(cmds, m.updateTodoForm(msg))
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		if m.currentPage != PageTodoDetails {
			return m, tea.Quit
		}

	case "tab":
		if m.focusedComponent == focusMenu {
			m.focusedComponent = focusContent
		} else {
			m.focusedComponent = focusMenu
		}
		return m, nil

	case "esc":
		if m.currentPage == PageTodoDetails {
			m.currentPage = PageMain
			m.focusedComponent = focusContent
			return m, nil
		}
		m.currentPage = PageMain
		m.focusedComponent = focusContent
		return m, nil

	case "t":
		if m.currentPage == PageThemes {
			m.currentPage = PageMain
		} else {
			m.currentPage = PageThemes
		}
		m.focusedComponent = focusContent
		return m, nil

	case "up", "k":
		if m.focusedComponent == focusMenu {
			if m.menuCursor > 0 {
				m.menuCursor--
			}
		}
		return m, nil

	case "down", "j":
		if m.focusedComponent == focusMenu {
			if m.menuCursor < len(m.menuItems)-1 {
				m.menuCursor++
			}
		}
		return m, nil

	case "enter":
		if m.focusedComponent == focusMenu {
			return m, m.menuItems[m.menuCursor].action(m)
		} else if m.currentPage == PageMain && m.focusedComponent == focusContent {
			// Open todo details
			if len(m.todos) > 0 && m.todoTable.Cursor() < len(m.todos) {
				m.selectedTodo = &m.todos[m.todoTable.Cursor()]
				m.loadTodoIntoForm()
				m.currentPage = PageTodoDetails
				m.focusedComponent = focusTodoForm
				return m, nil
			}
		} else if m.currentPage == PageThemes && m.focusedComponent == focusContent {
			// Apply selected theme
			selectedItem := m.themeList.SelectedItem()
			if item, ok := selectedItem.(themeListItem); ok {
				return m, applyTheme(m.services, m.themeManager, item.name)
			}
		}

	case " ":
		if m.currentPage == PageMain && m.focusedComponent == focusContent {
			// Toggle todo done status
			if len(m.todos) > 0 && m.todoTable.Cursor() < len(m.todos) {
				return m, toggleTodoDone(m.services, &m.todos[m.todoTable.Cursor()])
			}
		}
	}

	return m, nil
}

func (m *Model) handleMouseClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Handle clicks on the todo table checkbox column
	if m.currentPage == PageMain && msg.X <= 5 {
		// Calculate which row was clicked
		rowClicked := msg.Y - 5 // Adjust for header and borders
		if rowClicked >= 0 && rowClicked < len(m.todos) {
			return m, toggleTodoDone(m.services, &m.todos[rowClicked])
		}
	}
	return m, nil
}

func (m *Model) updateTodoForm(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+s":
			return m.saveTodoFromForm()
		}
	}

	m.nameInput, cmd = m.nameInput.Update(msg)
	cmds = append(cmds, cmd)

	m.descriptionInput, cmd = m.descriptionInput.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

// ============================================================================
// View (Rendering)
// ============================================================================

func (m *Model) View() string {
	theme := m.themeManager.GetCurrentTheme()

	// Create styles
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Primary.String())).
		Bold(true).
		Align(lipgloss.Center).
		Padding(1, 0)

	menuStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Primary.String())).
		Padding(1, 2).
		Width(28)

	contentStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Primary.String())).
		Padding(1, 2)

	// Render header
	header := headerStyle.Render(m.themeManager.CreateHeader())

	// Render menu
	menu := m.renderMenu()
	styledMenu := menuStyle.Render(menu)

	// Render content based on current page
	var content string
	switch m.currentPage {
	case PageMain:
		content = m.renderTodoTable()
	case PageThemes:
		content = m.themeList.View()
	case PageTodoDetails:
		content = m.renderTodoDetailsForm()
	case PageWelcome:
		content = "Welcome to TUIDOO! Your tasks will appear here"
	}

	styledContent := contentStyle.Render(content)

	// Layout
	mainLayout := lipgloss.JoinHorizontal(
		lipgloss.Top,
		styledMenu,
		styledContent,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		mainLayout,
	)
}

func (m *Model) renderMenu() string {
	var b strings.Builder
	theme := m.themeManager.GetCurrentTheme()

	for i, item := range m.menuItems {
		cursor := " "
		if m.focusedComponent == focusMenu && i == m.menuCursor {
			cursor = ">"
			// Highlight selected item
			style := lipgloss.NewStyle().
				Foreground(lipgloss.Color(theme.Colors.Primary.String())).
				Bold(true)
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, style.Render(item.label)))
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, item.label))
		}
	}

	return b.String()
}

func (m *Model) renderTodoTable() string {
	if len(m.todos) == 0 {
		return "No todos yet - Press 'n' to create a new task"
	}

	m.applyTableTheme()
	return m.todoTable.View()
}

func (m *Model) renderTodoDetailsForm() string {
	if m.selectedTodo == nil {
		return "No todo selected"
	}

	var b strings.Builder
	theme := m.themeManager.GetCurrentTheme()

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Primary.String())).
		Bold(true)

	b.WriteString(titleStyle.Render(fmt.Sprintf("Edit Todo: %s", m.selectedTodo.Name)))
	b.WriteString("\n\n")

	b.WriteString("Task Name:\n")
	b.WriteString(m.nameInput.View())
	b.WriteString("\n\n")

	b.WriteString("Description:\n")
	b.WriteString(m.descriptionInput.View())
	b.WriteString("\n\n")

	b.WriteString(fmt.Sprintf("Priority: %s\n", m.selectedTodo.Priority.String()))
	b.WriteString(fmt.Sprintf("Status: %s\n", m.selectedTodo.Status.String()))
	b.WriteString(fmt.Sprintf("Completed: %v\n", m.selectedTodo.Done))

	b.WriteString("\n")
	b.WriteString("Ctrl+S: Save | Esc: Cancel")

	return b.String()
}

// ============================================================================
// Helper Methods
// ============================================================================

func (m *Model) updateTableRows() {
	rows := make([]table.Row, 0, len(m.todos))

	for _, todo := range m.todos {
		rows = append(rows, table.Row{
			getStatusIcon(todo.Done),
			todo.Priority.String(),
			todo.Name,
			todo.Project.Name,
			todo.ToDoList.Name,
			todo.Status.String(),
		})
	}

	m.todoTable.SetRows(rows)
}

func (m *Model) applyTableTheme() {
	theme := m.themeManager.GetCurrentTheme()

	tableStyle := table.DefaultStyles()
	tableStyle.Header = tableStyle.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Primary.String())).
		BorderBottom(true).
		Bold(true)

	tableStyle.Selected = tableStyle.Selected.
		Foreground(lipgloss.Color(theme.Colors.Background.String())).
		Background(lipgloss.Color(theme.Colors.Primary.String())).
		Bold(false)

	m.todoTable.SetStyles(tableStyle)
}

func (m *Model) applyThemeStyles() {
	m.updateTableRows()
	m.applyTableTheme()
}

func (m *Model) loadTodoIntoForm() {
	if m.selectedTodo == nil {
		return
	}

	m.nameInput.SetValue(m.selectedTodo.Name)

	desc := getDescription(m.selectedTodo)
	m.descriptionInput.SetValue(desc)

	m.priorityIndex = getPriorityIndex(m.selectedTodo.Priority.String())
	m.statusIndex = getStatusIndex(m.selectedTodo.Status.String())
	m.doneCheckbox = m.selectedTodo.Done

	m.nameInput.Focus()
}

func (m *Model) saveTodoFromForm() tea.Cmd {
	if m.selectedTodo == nil {
		return nil
	}

	m.selectedTodo.Name = m.nameInput.Value()
	desc := m.descriptionInput.Value()
	m.selectedTodo.Description = &desc

	return saveTodo(m.services, m.selectedTodo)
}

// ============================================================================
// Commands (Async Operations)
// ============================================================================

func loadTodos(sc *services.ServiceCollection) tea.Cmd {
	return func() tea.Msg {
		todos, err := sc.ToDoService.GetAll(true)
		if err != nil {
			return errorMsg{err}
		}
		return todosLoadedMsg{todos}
	}
}

func toggleTodoDone(sc *services.ServiceCollection, todo *entities.ToDo) tea.Cmd {
	return func() tea.Msg {
		var err error
		if !todo.Done {
			err = sc.ToDoService.MarkAsComplete(todo.ID)
		} else {
			todo.Done = false
			err = sc.ToDoService.MarkAsIncomplete(todo.ID)
		}

		if err != nil {
			return errorMsg{err}
		}

		return todosLoadedMsg{} // Trigger reload
	}
}

func saveTodo(sc *services.ServiceCollection, todo *entities.ToDo) tea.Cmd {
	return func() tea.Msg {
		if err := sc.ToDoService.Update(todo); err != nil {
			return errorMsg{err}
		}
		return todoUpdatedMsg{todo}
	}
}

func applyTheme(sc *services.ServiceCollection, tm *managers.ThemeManager, themeName string) tea.Cmd {
	return func() tea.Msg {
		if err := tm.SetTheme(themeName); err != nil {
			return errorMsg{err}
		}

		if err := sc.SettingsService.SetActiveTheme(themeName); err != nil {
			return errorMsg{err}
		}

		log.Printf("✓ Theme changed to: %s", themeName)
		return themeChangedMsg{themeName}
	}
}

func showNotImplemented(feature string) tea.Cmd {
	return func() tea.Msg {
		log.Printf("%s feature coming soon!", feature)
		return nil
	}
}

// ============================================================================
// Theme List Item (for bubbles/list)
// ============================================================================

type themeListItem struct {
	name      string
	isCurrent bool
}

func (i themeListItem) FilterValue() string { return i.name }
func (i themeListItem) Title() string {
	if i.isCurrent {
		return "✓ " + i.name
	}
	return i.name
}
func (i themeListItem) Description() string { return "" }

type themeListDelegate struct{}

func (d themeListDelegate) Height() int                               { return 1 }
func (d themeListDelegate) Spacing() int                              { return 0 }
func (d themeListDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d themeListDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(themeListItem)
	if !ok {
		return
	}

	str := i.Title()
	if index == m.Index() {
		str = "> " + str
	} else {
		str = "  " + str
	}

	fmt.Fprint(w, str)
}

// ============================================================================
// Utility Functions
// ============================================================================

func getStatusIcon(done bool) string {
	if done {
		return "[✓]"
	}
	return "[ ]"
}

func getDescription(todo *entities.ToDo) string {
	if todo.Description != nil && *todo.Description != "" {
		return *todo.Description
	}
	if todo.Details != nil && *todo.Details != "" {
		return *todo.Details
	}
	return "No description available"
}

func getPriorityIndex(priority string) int {
	switch priority {
	case "Low":
		return 0
	case "Medium":
		return 1
	case "High":
		return 2
	default:
		return 0
	}
}

func getStatusIndex(status string) int {
	for i, s := range enums.StatusOptions {
		if s == status {
			return i
		}
	}
	return 0
}

func applyUserTheme(sc *services.ServiceCollection, tm *managers.ThemeManager) error {
	settings, err := sc.SettingsService.GetAllSettings()
	if err != nil {
		return fmt.Errorf("failed to get settings: %w", err)
	}

	if err := tm.SetTheme(settings.ActiveThemeID); err != nil {
		return fmt.Errorf("failed to set theme: %w", err)
	}

	log.Printf("✓ Applied theme: %s", settings.ActiveThemeID)
	return nil
}

// SeedDatabase seeds the database with sample data
func SeedDatabase() {
	sc, err := services.NewServiceCollection()
	if err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer sc.Close()

	if err := services.Seed(sc.DbService); err != nil {
		log.Fatalf("❌ Seeding failed: %v", err)
	}

	fmt.Println("✅ Database seeded successfully")
	os.Exit(0)
}

// ResetDatabase cleans and reseeds the database
func ResetDatabase() {
	sc, err := services.NewServiceCollection()
	if err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer sc.Close()

	fmt.Println("⚠️  This will delete all existing data. Continue? (y/N)")
	var response string
	fmt.Scanln(&response)

	if response != "y" && response != "Y" {
		fmt.Println("Reset cancelled")
		os.Exit(0)
	}

	if err := services.ResetAndSeed(sc.DbService); err != nil {
		log.Fatalf("❌ Reset failed: %v", err)
	}

	fmt.Println("✅ Database reset successfully")
	os.Exit(0)
}

// CleanDatabase removes all data from the database
func CleanDatabase() {
	sc, err := services.NewServiceCollection()
	if err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer sc.Close()

	fmt.Println("⚠️  This will delete all data. Continue? (y/N)")
	var response string
	fmt.Scanln(&response)

	if response != "y" && response != "Y" {
		fmt.Println("Clean cancelled")
		os.Exit(0)
	}

	if err := services.CleanDatabase(sc.DbService); err != nil {
		log.Fatalf("❌ Clean failed: %v", err)
	}

	fmt.Println("✅ Database cleaned successfully")
	os.Exit(0)
}
