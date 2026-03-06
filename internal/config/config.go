package config

import (
	"os"
	"path/filepath"
)

// Config contains all the app configuration
type Config struct {
	HomeDir   string
	RepoURL   string
	LocalPath string
	Symlinks  []SymlinkMapping
}

// SymlinkMapping defines the mapping between source and destination files
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
		HomeDir:   homeDir,
		RepoURL:   "https://github.com/lexusletz/dotfiles.git",
		LocalPath: filepath.Join(homeDir, "dotfiles"),
		Symlinks:  getDefaultSymlinks(),
	}, nil
}

// getDefaultSymlinks returns the default Symlinks configuration
func getDefaultSymlinks() []SymlinkMapping {
	return []SymlinkMapping{
		{Source: "nvim", Target: ".config/nvim"},
		{Source: "alacritty", Target: ".config/alacritty"},
		{Source: "tmux/.tmux.conf", Target: ".tmux.conf"},
		{Source: "ghostty", Target: ".config/ghostty"},
	}
}
