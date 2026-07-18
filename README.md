# mandrill

> **Disclaimer:** This project is named "mandrill" and is a package manager / git TUI tool.
> It is **NOT** affiliated with endorsed by, or related to [Mailchimp's Mandrill](https://mandrillapp.com/)
> transactional email service or any of its associated libraries, wrappers, or API clients.

A package manager and git TUI on steroids, built with Go and Rust.

## Features

- Install, update, and remove CLI tools from a git-based registry
- Interactive git TUI (like lazygit, but more powerful)
- GitHub CLI (`gh`) integration
- 13 built-in color themes (catppuccin, dracula, tokyo-night, etc.)
- Idle screensaver with animated ASCII art
- Cross-platform: macOS, Linux, Windows

## Installation

### macOS

**Quick install:**
```bash
curl -sL https://raw.githubusercontent.com/838288383838383/mandrill/master/install.sh | bash
```

**Homebrew:**
```bash
brew tap 838288383838383/mandrill https://github.com/838288383838383/mandrill-homebrew.git
brew install mandrill
```

**Go:**
```bash
GOPROXY=direct go install github.com/838288383838383/mandrill/cmd/mandrill@latest
```

### Linux

**Binary download:**
```bash
curl -sL https://github.com/838288383838383/mandrill/releases/latest/download/mandrill-linux-amd64 -o /usr/local/bin/mandrill
chmod +x /usr/local/bin/mandrill
```

**Go:**
```bash
GOPROXY=direct go install github.com/838288383838383/mandrill/cmd/mandrill@latest
```

### Windows

**Quick install (PowerShell):**
```powershell
Invoke-WebRequest -Uri "https://github.com/838288383838383/mandrill/releases/latest/download/mandrill-windows-amd64.exe" -OutFile "$env:LOCALAPPDATA\mandrill\bin\mandrill.exe"
$env:PATH += ";$env:LOCALAPPDATA\mandrill\bin"
```

**Batch install:**
```
install.bat
```

**Go:**
```
GOPROXY=direct go install github.com/838288383838383/mandrill/cmd/mandrill@latest
```

**Scoop:**
```
scoop install mandrill
```

## Usage

```bash
mandrill install <tool>     # Install a tool
mandrill uninstall <tool>   # Remove a tool
mandrill list               # List installed tools
mandrill search <query>     # Search registry
mandrill git                # Launch git TUI
mandrill theme list         # List available themes
mandrill auth status        # Check GitHub pairing
```

## First Run

On first launch, mandrill will prompt you to set up GitHub CLI integration:

```
? Pair with GitHub CLI (gh)? [y/n/fuckoff]
```

## Platform Paths

| OS | Config Location |
|----|-----------------|
| macOS | `~/Library/Application Support/mandrill/` |
| Linux | `~/.mandrill/` |
| Windows | `%LOCALAPPDATA%\mandrill\` |

## Build from Source

```bash
git clone https://github.com/838288383838383/mandrill.git
cd mandrill
make build          # Build for current platform
make build-all      # Build for all platforms
make mac-sign       # Sign macOS binaries
make release        # Build + GitHub release
```

## License

MIT License - see [LICENSE](LICENSE) for details.
