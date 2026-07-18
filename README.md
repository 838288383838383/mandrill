# mandrill

> **⚠️ Disclaimer:** This project is named "mandrill" and is a package manager / git TUI tool.
> It is **NOT** affiliated with, endorsed by, or related to [Mailchimp's Mandrill](https://mandrillapp.com/)
> transactional email service or any of its associated libraries, wrappers, or API clients.

A package manager and git TUI on steroids, built with Go and Rust.

## Features

- Install, update, and remove CLI tools from a git-based registry
- Interactive git TUI (like lazygit, but more powerful)
- GitHub CLI (`gh`) integration
- 13 built-in color themes (catppuccin, dracula, tokyo-night, etc.)
- Idle screensaver with animated ASCII art

## Installation

```bash
# Clone and build
git clone https://github.com/oddman/mandrill.git
cd mandrill
go build ./cmd/mandrill/
```

## Usage

```bash
mandrill install <tool>     # Install a tool
mandrill uninstall <tool>   # Remove a tool
mandrill list               # List installed tools
mandrill search <query>     # Search registry
mandrill git                # Launch git TUI
mandrill theme list         # List available themes
```

## First Run

On first launch, mandrill will prompt you to set up GitHub CLI integration:

```
? Pair with GitHub CLI (gh)? [y/n/fuckoff]
```

## License

MIT License - see [LICENSE](LICENSE) for details.
