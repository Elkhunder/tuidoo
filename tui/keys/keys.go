package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	// Navigation
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	PageUp   key.Binding
	PageDown key.Binding
	Home     key.Binding
	End      key.Binding

	// Actions
	Enter  key.Binding
	Space  key.Binding
	Escape key.Binding
	Tab    key.Binding

	// Application
	Quit    key.Binding
	Refresh key.Binding
	Help    key.Binding

	// Todo specific
	NewTodo    key.Binding
	EditTodo   key.Binding
	DeleteTodo key.Binding
	ToggleDone key.Binding

	// Views
	ToggleThemes key.Binding
	ViewProjects key.Binding
}

var Keys = &KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "right"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("pgup", "b"),
		key.WithHelp("pgup", "page up"),
	),
	PageDown: key.NewBinding(
		key.WithKeys("pgdown", "f"),
		key.WithHelp("pgdown", "page down"),
	),
	Home: key.NewBinding(
		key.WithKeys("home", "g"),
		key.WithHelp("g/home", "go to start"),
	),
	End: key.NewBinding(
		key.WithKeys("end", "G"),
		key.WithHelp("G/end", "go to end"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Space: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle done"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel/back"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch focus"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	NewTodo: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new task"),
	),
	EditTodo: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit task"),
	),
	DeleteTodo: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete task"),
	),
	ToggleDone: key.NewBinding(
		key.WithKeys(" ", "x"),
		key.WithHelp("space/x", "toggle done"),
	),
	ToggleThemes: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "themes"),
	),
	ViewProjects: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "projects"),
	),
}
