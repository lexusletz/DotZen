# DotZen

**A modern and elegant CLI tool for managing your dotfiles with ease**

DotZen automates the synchronization of your dotfiles repository and creates symlinks automatically, keeping your configuration synchronized across all your systems.

## Features

- **Automatic Synchronization** - Clone or update your dotfiles repository
- **Symlink Management** - Automatically create symbolic links to your configuration files
- **Automatic Backup** - Backup existing files before creating symlinks
- **Cross-Platform** - Compatible with macOS, Linux, and Windows
- **Fast and Efficient** - Written in Go for maximum performance
- **Simple Configuration** - Minimal setup required

## Installation

### Download Pre-built Releases

Download the precompiled binary for your system from the [Releases](https://github.com/lexusletz/dotzen/releases) page:

#### macOS
```bash
# Intel Macs
curl -L https://github.com/lexusletz/dotzen/releases/latest/download/dotzen-darwin-amd64.tar.gz | tar -xz
sudo mv dotzen /usr/local/bin/

# Apple Silicon (M1/M2)
curl -L https://github.com/lexusletz/dotzen/releases/latest/download/dotzen-darwin-arm64.tar.gz | tar -xz
sudo mv dotzen /usr/local/bin/
```

#### Linux
```bash
# x86_64
curl -L https://github.com/lexusletz/dotzen/releases/latest/download/dotzen-linux-amd64.tar.gz | tar -xz
sudo mv dotzen /usr/local/bin/

# ARM64
curl -L https://github.com/lexusletz/dotzen/releases/latest/download/dotzen-linux-arm64.tar.gz | tar -xz
sudo mv dotzen /usr/local/bin/
```

### Build from Source

#### Requirements
- [Go](https://golang.org/dl/) 1.21 or higher
- Git

#### Installation
```bash
# Clone the repository
git clone https://github.com/lexusletz/dotzen.git
cd dotzen

# Install Unix/MacOS
./build --install
```

## Usage

### Basic Usage
```bash
# Run DotZen
dotzen
```

This will:
1. Verify Git is installed
2. Clone or update your dotfiles repository
3. Create symlinks for all configured files
4. Automatically backup existing files

### Configuration

DotZen looks for your dotfiles repository at: `https://github.com/lexusletz/dotfiles.git`

#### Customize Configuration

Edit the `internal/config/config.go` file to customize:

```go
// In the New() function, change:
RepoURL: "https://github.com/YOUR-USERNAME/dotfiles.git",

// In getDefaultSymlinks(), add your files:
{Source: "nvim", Target: ".config/nvim"},
{Source: "alacritty", Target: ".config/alacritty"},
// ... more configurations
```

### Recommended Dotfiles Structure
```
dotfiles/
├── .vimrc
├── .zshrc
├── .gitconfig
├── .tmux.conf
├── nvim/
│   ├── init.vim
│   └── ...
├── alacritty/
│   └── alacritty.yml
└── README.md
```

## Development

### Project Structure
```
dotzen/
├── cmd/dotzen/          # Main entry point
├── internal/
│   ├── config/          # Configuration
│   ├── git/             # Git operations
│   ├── symlink/         # Symlink management
│   └── dotfiles/        # Main orchestration
├── bin/                 # Compiled binaries
├── dist/                # Release archives
├── Makefile             # Build targets
├── build.sh             # Unix build script
├── build.ps1            # Windows build script
└── README.md
```

## Advanced Configuration

### Custom Symlinks

Symlinks are configured in `internal/config/config.go`:

```go
type SymlinkMapping struct {
    Source string // File in the dotfiles repo
    Target string // Destination on the system (relative to $HOME)
}
```

Examples:
```go
{ Source: ".vimrc", Target: ".vimrc" },                                 // Simple file
{ Source: "nvim", Target: ".config/nvim" },                             // Full directory
{ Source: "scripts/my-script.sh", Target: ".local/bin/my-script.sh" },  // Subdirectory
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `HOME` | User home directory | Auto-detected |

## Contributing

Contributions are welcome! Please:

1. Fork the project
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Reporting Bugs

Use [GitHub Issues](https://github.com/lexusletz/dotzen/issues) to report bugs or request features.

## Changelog

### v1.0.0
- First stable release
- Complete symlink management
- Automatic backup of existing files
- Cross-platform support
- Cross-compilation for all platforms

## Security

- DotZen creates automatic backups before overwriting files
- Only modifies files in your home directory
- Does not require administrator permissions (except for global installation)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by [GNU Stow](https://www.gnu.org/software/stow/) and other dotfiles managers
- Built with ❤️ using [Go](https://golang.org/)

## Support

- Email: [jordypinosdev@gmail.com](mailto:jordypinosdev@gmail.com)
- Issues: [GitHub Issues](https://github.com/lexusletz/dotzen/issues)
- Discussions: [GitHub Discussions](https://github.com/lexusletz/dotzen/discussions)

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/lexusletz">@lexusletz</a>
</p>

<p align="center">
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Made%20with-Go-00ADD8?style=for-the-badge&logo=go" alt="Made with Go">
  </a>
  <a href="https://github.com/lexusletz/dotzen/blob/main/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" alt="License: MIT">
  </a>
  <a href="https://github.com/lexusletz/dotzen/releases">
    <img src="https://img.shields.io/github/v/release/lexusletz/dotzen?style=for-the-badge" alt="Latest Release">
  </a>
</p>
