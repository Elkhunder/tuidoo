package main

import (
	"fmt"
	"log"
	"tuidoo/entities"

	"tuidoo/managers"
	"tuidoo/services"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/thiagokokada/dark-mode-go"
)

type App struct {
	tviewApp     *tview.Application
	services     *services.ServiceCollection
	themeManager *managers.ThemeManager

	// UI Components
	header        *tview.TextView
	menu          *tview.List
	todoList      *tview.List
	themeSelector *tview.List
	mainContent   *tview.TextView
	pages         *tview.Pages
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run() error {
	// Initialize application
	app, err := initializeApp()
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}
	defer app.cleanup()

	// Build UI
	if err := app.buildUI(); err != nil {
		return fmt.Errorf("failed to build UI: %w", err)
	}

	// Run application
	log.Println("🚀 Starting TUIDOO...")
	if err := app.tviewApp.SetRoot(app.createLayout(), true).EnableMouse(true).Run(); err != nil {
		return fmt.Errorf("error running tuidoo: %w", err)
	}

	return nil
}

func initializeApp() (*App, error) {
	log.Println("Initializing application...")

	// Check system theme
	if err := checkSystemTheme(); err != nil {
		log.Printf("Warning: Could not detect system theme: %v", err)
	}

	// Initialize services
	sc, err := services.NewServiceCollection()
	if err != nil {
		return nil, fmt.Errorf("service initialization failed: %w", err)
	}

	// Load data
	projects, toDoLists, todos, err := loadData(sc)
	if err != nil {
		sc.Close()
		return nil, fmt.Errorf("data loading failed: %w", err)
	}

	log.Printf("✓ Loaded %d projects, %d lists, %d todos", len(projects), len(toDoLists), len(todos))

	// Initialize theme manager
	themeManager := managers.NewThemeManager()
	if err := applyUserTheme(sc, themeManager); err != nil {
		sc.Close()
		return nil, fmt.Errorf("theme initialization failed: %w", err)
	}

	return &App{
		tviewApp:     tview.NewApplication(),
		services:     sc,
		themeManager: themeManager,
	}, nil
}

func checkSystemTheme() error {
	isDark, err := dark.IsDarkMode()
	if err != nil {
		return err
	}

	if isDark {
		log.Println("System is using dark mode")
	} else {
		log.Println("System is using light mode")
	}

	return nil
}

func loadData(sc *services.ServiceCollection) ([]entities.Project, []entities.ToDoList, []entities.ToDo, error) {
	projectList, err := sc.ProjectService.GetAll(true)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to fetch projects: %w", err)
	}

	listList, err := sc.ToDoListService.GetAll(true)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to fetch todo lists: %w", err)
	}

	todoList, err := sc.ToDoService.GetAll(true)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to fetch todos: %w", err)
	}

	return projectList, listList, todoList, nil
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

func (app *App) cleanup() {
	log.Println("Shutting down...")
	if err := app.services.Close(); err != nil {
		log.Printf("Error during cleanup: %v", err)
	}
}

func (app *App) buildUI() error {
	log.Println("Building UI...")

	app.createHeader()
	app.createMenu()
	app.createTodoList()
	app.createThemeSelector()
	app.createMainContent()
	app.createPages()
	app.setupInputCapture()

	log.Println("✓ UI built successfully")
	return nil
}

func (app *App) createHeader() {
	theme := app.themeManager.GetCurrentTheme()

	app.header = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(app.themeManager.CreateHeader())

	app.header.SetBackgroundColor(theme.Colors.Background)
	app.header.SetBorderPadding(2, 2, 0, 0)
}

func (app *App) createMenu() {
	app.menu = tview.NewList().SetHighlightFullLine(true)

	app.menu.AddItem("➕ New Task", "Create new task", 'n', func() {
		// TODO: Implement new task form
		app.showNotImplemented("New Task")
	}).
		AddItem("⚙️  Settings", "App settings", 's', func() {
			app.pages.SwitchToPage("themes")
			app.tviewApp.SetFocus(app.themeSelector)
		}).
		AddItem("❌ Quit", "Exit application", 'q', func() {
			app.tviewApp.Stop()
		})

	app.menu.SetBorder(true).SetTitle(" Menu ")
}

func (app *App) createTodoList() {
	app.todoList = tview.NewList()

	// Fetch and populate todos
	todos, err := app.services.ToDoService.GetAll(true)
	if err != nil {
		log.Printf("Failed to load todos: %v", err)
		return
	}

	for _, todo := range todos {
		todoItem := todo // Capture for closure
		status := "[ ]"
		if todoItem.Done {
			status = "[✓]"
		}

		displayText := fmt.Sprintf("%s %s", status, todoItem.Name)
		secondaryText := fmt.Sprintf("Status: %s | Priority: %s",
			todoItem.Status.String(),
			todoItem.Priority.String())

		app.todoList.AddItem(displayText, secondaryText, 0, func() {
			app.showTodoDetails(&todoItem)
		})
	}

	app.todoList.SetBorder(true).SetTitle(" tuidoo — pre-alpha ")
}

func (app *App) createThemeSelector() {
	app.themeSelector = tview.NewList()

	for _, themeName := range app.themeManager.GetThemeNames() {
		name := themeName // Capture for closure

		app.themeSelector.AddItem(name, "", 0, func() {
			if err := app.applyTheme(name); err != nil {
				log.Printf("Failed to apply theme: %v", err)
			}
		})
	}

	app.themeSelector.SetBorder(true).SetTitle(" Select Theme ")
}

func (app *App) createMainContent() {
	app.mainContent = tview.NewTextView().
		SetText("[-]Welcome to TUIDOO!\n\nYour tasks will appear here.\n\nPress 't' to toggle theme selector.\nPress 'Tab' to switch focus.").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	app.mainContent.SetBorder(true).SetTitle(" Main ")
}

func (app *App) createPages() {
	app.pages = tview.NewPages().
		AddPage("main", app.todoList, true, true).
		AddPage("themes", app.themeSelector, true, false).
		AddPage("content", app.mainContent, true, false)
}

func (app *App) createLayout() *tview.Flex {
	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(app.menu, 30, 0, true).
		AddItem(app.pages, 0, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(app.header, 10, 0, false).
		AddItem(mainLayout, 0, 1, true)

	return layout
}

func (app *App) setupInputCapture() {
	app.tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			app.handleTabKey()
		case tcell.KeyEscape:
			app.pages.SwitchToPage("main")
			app.tviewApp.SetFocus(app.todoList)
		}

		switch event.Rune() {
		case 't':
			app.toggleThemeSelector()
		case 'q':
			app.tviewApp.Stop()
		}

		return event
	})
}

func (app *App) handleTabKey() {
	switch app.tviewApp.GetFocus() {
	case app.menu:
		// Get current page and focus it
		frontPage, _ := app.pages.GetFrontPage()
		switch frontPage {
		case "main":
			app.tviewApp.SetFocus(app.todoList)
		case "themes":
			app.tviewApp.SetFocus(app.themeSelector)
		default:
			app.tviewApp.SetFocus(app.mainContent)
		}
	default:
		app.tviewApp.SetFocus(app.menu)
	}
}

func (app *App) toggleThemeSelector() {
	frontPage, _ := app.pages.GetFrontPage()

	if frontPage == "themes" {
		app.pages.SwitchToPage("main")
		app.tviewApp.SetFocus(app.todoList)
	} else {
		app.pages.SwitchToPage("themes")
		app.tviewApp.SetFocus(app.themeSelector)
	}
}

func (app *App) applyTheme(themeName string) error {
	if err := app.themeManager.SetTheme(themeName); err != nil {
		return fmt.Errorf("failed to set theme: %w", err)
	}

	// Save theme preference
	if err := app.services.SettingsService.SetActiveTheme(themeName); err != nil {
		return fmt.Errorf("failed to save theme preference: %w", err)
	}

	// Update UI
	theme := app.themeManager.GetCurrentTheme()
	app.header.SetText(app.themeManager.CreateHeader())
	app.header.SetBackgroundColor(theme.Colors.Background)

	log.Printf("✓ Theme changed to: %s", themeName)
	return nil
}

func (app *App) showTodoDetails(todo interface{}) {
	// TODO: Implement todo details view
	app.showNotImplemented("Todo Details")
}

func (app *App) showNotImplemented(feature string) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("%s feature coming soon!", feature)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.pages.SwitchToPage("main")
			app.tviewApp.SetFocus(app.todoList)
		})

	app.pages.AddPage("modal", modal, true, true)
}
