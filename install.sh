#!/bin/bash
set -e

echo "8b    d8    db    88b 88 8888b.  88\"\"Yb 88 88     88"
echo "88b  d88   dPYb   88Yb88  8I  Yb 88__dP 88 88     88"
echo "88YbdP88  dP__Yb  88 Y88  8I  dY 88\"Yb  88 88  .o 88  .o"
echo "88 YY 88 dP\"\"\"\"Yb 88  Y8 8888Y\"  88  Yb 88 88ood8 88ood8"
echo ""
echo "  Package manager & git TUI — macOS installer"
echo ""

ARCH=$(uname -m)
if [ "$ARCH" = "arm64" ]; then
    BINARY="mandrill-darwin-arm64"
    echo "Detected: Apple Silicon (arm64)"
elif [ "$ARCH" = "x86_64" ]; then
    BINARY="mandrill-darwin-amd64"
    echo "Detected: Intel (x86_64)"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

INSTALL_DIR="$HOME/.mandrill/bin"
mkdir -p "$INSTALL_DIR"

echo "Downloading $BINARY..."
curl -sL "https://github.com/838288383838383/mandrill/releases/latest/download/$BINARY" -o "$INSTALL_DIR/mandrill"
chmod +x "$INSTALL_DIR/mandrill"

echo "Installed to $INSTALL_DIR/mandrill"

SHELL_RC=""
if [ -n "$ZSH_VERSION" ] || [ -f "$HOME/.zshrc" ]; then
    SHELL_RC="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ] || [ -f "$HOME/.bashrc" ]; then
    SHELL_RC="$HOME/.bashrc"
fi

if [ -n "$SHELL_RC" ]; then
    if ! grep -q ".mandrill/bin" "$SHELL_RC" 2>/dev/null; then
        echo "" >> "$SHELL_RC"
        echo "# mandrill" >> "$SHELL_RC"
        echo "export PATH=\"\$HOME/.mandrill/bin:\$PATH\"" >> "$SHELL_RC"
        echo "Added PATH to $SHELL_RC"
        echo "Run: source $SHELL_RC"
    fi
fi

echo ""
echo "Done! Run 'mandrill --help' to get started."
