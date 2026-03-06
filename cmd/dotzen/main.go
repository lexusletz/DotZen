package main

import (
	"dotzen/internal/config"
	"dotzen/internal/dotfiles"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		fmt.Printf("Error obtaining configuration: %v\n", err)
		os.Exit(1)
	}

	manager := dotfiles.New(cfg)

	if err := manager.Setup(); err != nil {
		fmt.Printf("Error in setup: %v\n", err)
		os.Exit(1)
	}
}
