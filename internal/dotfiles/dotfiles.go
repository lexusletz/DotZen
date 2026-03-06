package dotfiles

import (
	"fmt"

	"dotzen/internal/config"
	"dotzen/internal/git"
	"dotzen/internal/symlink"
)

// Manager manages all dotfiles operations
type Manager struct {
	config    *config.Config
	repo      *git.Repository
	symlinker *symlink.Manager
}

// New creates a new instance of Manager
func New(cfg *config.Config) *Manager {
	repo := git.New(cfg.RepoURL, cfg.LocalPath)
	symlinker := symlink.New(cfg.HomeDir, cfg.LocalPath, cfg.Symlinks)

	return &Manager{
		config:    cfg,
		repo:      repo,
		symlinker: symlinker,
	}
}

// Setup execute the entire configuration process
func (m *Manager) Setup() error {
	fmt.Println("DotZen working correctly!")

	if err := git.IsGitInstalled(); err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println("Git found")

	fmt.Println("Your home directory is: ", m.config.HomeDir)
	fmt.Printf("Local path: %s\n", m.config.LocalPath)
	fmt.Printf("Remote path: %s\n", m.config.RepoURL)

	if err := m.repo.Sync(); err != nil {
		return fmt.Errorf("%v", err)
	}

	if err := m.symlinker.CreateAll(); err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Println("\nDotZen successfully completed!")
	return nil
}
