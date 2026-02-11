package config

import (
	"os"
	"path/filepath"
)

// Config contiene toda la configuración de la aplicación
type Config struct {
	HomeDir 	string
	RepoURL 	string
	LocalPath 	string
	Symlinks 	[]SymlinkMapping
}

// SymlinkMapping define el mapeo entre archivos fuente y destino
type SymlinkMapping struct {
	Source string
	Target string
}

func New() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	return &Config{
		HomeDir: homeDir,
		RepoURL: "https://github.com/lexusletz/dotfiles.git",
		LocalPath: filepath.Join(homeDir, "dotfiles"),
		Symlinks: getDefaultSymlinks(),
	}, nil
}

// getDefaultSymlinks retorna la configuración por defecto de Symlinks
func getDefaultSymlinks() []SymlinkMapping {
	return []SymlinkMapping {
		{Source: "nvim", Target: ".config/nvim"},
		{Source: "alacritty", Target: ".config/alacritty"},
		{Source: "tmux/.tmux.conf", Target: ".tmux.conf"},
		{Source: "ghostty", Target: ".config/ghostty"},
	}
}
