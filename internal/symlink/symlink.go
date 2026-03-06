package symlink

import (
	"dotzen/internal/config"
	"fmt"
	"os"
	"path/filepath"
)

// Manager handles the creation and management of Symlinks
type Manager struct {
	homeDir      string
	dotfilesPath string
	mappings     []config.SymlinkMapping
}

// New creates a new instance of Manager
func New(homeDir, dotfilesPath string, mappings []config.SymlinkMapping) *Manager {
	return &Manager{
		homeDir:      homeDir,
		dotfilesPath: dotfilesPath,
		mappings:     mappings,
	}
}

// CreateAll creates all the Symlinks configured
func (m *Manager) CreateAll() error {
	fmt.Println("\nCreating symlinks...")

	for _, mapping := range m.mappings {
		if err := m.createSingle(mapping); err != nil {
			fmt.Printf("Error creating Symlink for %s: %v\n", mapping.Target, err)
		}
	}

	fmt.Println("Symlink creation completed")
	return nil
}

// createSingle create a single Symlink
func (m *Manager) createSingle(mapping config.SymlinkMapping) error {
	sourcePath := filepath.Join(m.dotfilesPath, mapping.Source)
	targetPath := filepath.Join(m.homeDir, mapping.Target)

	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		fmt.Printf("Source file not found, skipping: %s\n", sourcePath)
		return nil
	}

	if _, err := os.Lstat(targetPath); err == nil {
		if err := m.handleExisting(sourcePath, targetPath); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("error creating directory %s: %v", filepath.Dir(targetPath), err)
	}

	if err := os.Symlink(sourcePath, targetPath); err != nil {
		return fmt.Errorf("error creating Symlink %s -> %s: %v", targetPath, sourcePath, err)
	}

	fmt.Printf("Symlink created: %s -> %s\n", targetPath, sourcePath)
	return nil
}

func (m *Manager) handleExisting(sourcePath, targetPath string) error {
	if link, err := os.Readlink(targetPath); err == nil {
		if link == sourcePath {
			fmt.Printf("Symlink already exists: %s -> %s\n", targetPath, sourcePath)
			return nil
		}

		fmt.Printf("Updating existing Symlink: %s\n", targetPath)
	} else {
		fmt.Printf("Source file found: %s (creating backup)\n", targetPath)

		backupPath := targetPath + ".backup"
		if err := os.Rename(targetPath, backupPath); err != nil {
			return fmt.Errorf("error creating backup of %s: %v", targetPath, err)
		}

		fmt.Printf("Backup created: %s\n", backupPath)
	}

	if err := os.Remove(targetPath); err != nil {
		return fmt.Errorf("error deleting %s: %v", targetPath, err)
	}

	return nil
}
