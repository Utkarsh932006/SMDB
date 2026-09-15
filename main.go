package main

import (
	"fmt"
	"os"

	"github.com/Utkarsh932006/SMDB/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("🎬 SMDB — Simple Movie DataBrowser")
	fmt.Printf("API token loaded (%d chars). Ready to build!\n", len(cfg.APIToken))
}
