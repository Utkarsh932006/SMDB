package main

import (
	"fmt"
	"os"

	"github.com/Utkarsh932006/SMDB/internal/api"
	"github.com/Utkarsh932006/SMDB/internal/config"
	"github.com/Utkarsh932006/SMDB/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	client := api.NewClient(cfg.APIToken)
	model := tui.NewModel(client)

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running SMDB: %v\n", err)
		os.Exit(1)
	}
}
