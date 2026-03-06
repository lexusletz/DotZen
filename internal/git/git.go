package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Repository manages the operations of the git repository
type Repository struct {
	URL       string
	LocalPath string
}

// New creates a new instance of Repository
func New(url, localPath string) *Repository {
	return &Repository{
		URL:       url,
		LocalPath: localPath,
	}
}

// IsGitInstalled checks if Git is available on the system
func IsGitInstalled() error {
	_, err := exec.LookPath("git")
	if err != nil {
		return fmt.Errorf("Git is not installed or is not in the PATH")
	}
	return nil
}

// Exists checks if the repository exists locally
func (r *Repository) Exists() bool {
	gitDir := filepath.Join(r.LocalPath, ".git")
	_, err := os.Stat(gitDir)
	return err == nil
}

// Clone clones the remote repo
func (r *Repository) Clone() error {
	fmt.Println("Cloning repository...")

	cmd := exec.Command("git", "clone", r.URL, r.LocalPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error cloning repository: %v", err)
	}

	fmt.Println("Repository cloned successfully")
	return nil
}

// Pull updates the local repository
func (r *Repository) Pull() error {
	fmt.Println("Repository found, updating...")

	cmd := exec.Command("git", "pull", "origin", "main")
	cmd.Dir = r.LocalPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("Repository updated successfully")
	return nil
}

// Sync clone or update the repository as needed
func (r *Repository) Sync() error {
	if r.Exists() {
		return r.Pull()
	}

	return r.Clone()
}
