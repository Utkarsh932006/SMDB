package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

// KeyMap defines the keybindings for the application.
type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	Enter    key.Binding
	Back     key.Binding
	Tab      key.Binding
	PrevTab  key.Binding
	Cat1     key.Binding
	Cat2     key.Binding
	Cat3     key.Binding
	Cat4     key.Binding
	Help     key.Binding
	Quit     key.Binding
}

// DefaultKeyMap returns the default set of keybindings matching the requirements.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "view details"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc", "backspace"),
			key.WithHelp("esc/backspace", "go back"),
		),
		Tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next category"),
		),
		PrevTab: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("shift+tab", "previous category"),
		),
		Cat1: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", "now playing"),
		),
		Cat2: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", "popular"),
		),
		Cat3: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("3", "top rated"),
		),
		Cat4: key.NewBinding(
			key.WithKeys("4"),
			key.WithHelp("4", "upcoming"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}
