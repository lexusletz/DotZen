# AGENTS.md - DotZen Development Guidelines

## Overview

DotZen is a Go CLI tool for managing dotfiles with symlink automation. This document provides guidelines for agentic coding agents working on this codebase.

## Build Commands

### Standard Build
```bash
make build                    # Compile for current platform
./build.sh                   # Unix build script
```

### Multi-Platform Build
```bash
make build-all               # Build for all platforms
./build.sh --all            # Unix: all platforms
```

### Release
```bash
make release                 # Create release archives
./build.sh --release        # Unix: create release
```

### Testing
```bash
make test                    # Run all tests
go test -v ./...            # Direct go test
go test ./internal/config   # Test single package
go test -v ./internal/git   # Test with verbose output
go test -run TestName       # Run specific test
```

### Other Commands
```bash
make clean                   # Clean build artifacts
make install                # Build and install to /usr/local/bin
./build.sh --install        # Unix install
make help                   # Show available commands
```

## Code Style Guidelines

### Imports
- Standard library first, then third-party, then local packages
- Use grouped imports with parentheses
- Local packages use `dotzen/internal/<package>` format
- Example:
```go
import (
    "fmt"
    "os"
    "path/filepath"
    "dotzen/internal/config"
)
```

### Naming Conventions
- **Packages**: lowercase, single word or short name describing functionality
- **Types**: PascalCase (e.g., `Config`, `Repository`, `Manager`)
- **Variables**: camelCase (e.g., `homeDir`, `localPath`)
- **Constants**: camelCase for unexported, PascalCase for exported
- **Functions**: camelCase for unexported, PascalCase for exported
- **Interfaces**: Add "er" suffix when appropriate (e.g., `Reader`)

### Error Handling
- Return errors from functions that can fail
- Use descriptive error messages with context
- Handle errors at the appropriate level
- Example pattern:
```go
if err != nil {
    return fmt.Errorf("error doing thing: %v", err)
}
```

### Function Organization
- Receiver methods for operations on types
- Place `New` constructor functions in each package
- Keep functions focused and under 50 lines when possible
- Use early returns to reduce nesting

### Comments
- Use comments for exported types and functions
- Comment format: "TypeName does X" or "FunctionName does Y"
- No comments on unexported code unless complex logic
- Example:
```go
// Config contains all application configuration
type Config struct { ... }

// New creates a new Config instance
func New() (*Config, error) { ... }
```

### Types and Structs
- Use structs for data containers
- Keep fields aligned with tabs
- Example:
```go
type SymlinkMapping struct {
    Source string
    Target string
}
```

### Formatting
- Run `go fmt` before committing
- Use tabs for indentation
- No trailing whitespace
- Blank line between import groups and code

### Project Structure
```
dotzen/
├── cmd/dotzen/          # Entry point
├── internal/
│   ├── config/          # Configuration
│   ├── git/            # Git operations
│   ├── symlink/        # Symlink management
│   └── dotfiles/       # Main orchestration
├── Makefile            # Build targets
└── build.sh            # Build script
```

### Code Patterns

#### Constructor Pattern
```go
func New(param1 string, param2 int) *Type {
    return &Type{
        field1: param1,
        field2: param2,
    }
}
```

#### Receiver Methods
```go
func (r *Repository) MethodName() error {
    // implementation
}
```

#### Error Handling in Main
```go
func main() {
    cfg, err := config.New()
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
        os.Exit(1)
    }
    // continue
}
```

## Development Workflow

1. Make changes to code
2. Run `make build` to verify compilation
3. Run `make test` to execute tests
4. Commit with descriptive message

## Key Dependencies

- Go 1.21+ required
- Git for dotfiles operations
- Standard library only (no external dependencies)
