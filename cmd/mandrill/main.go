package main

import (
	"fmt"
	"os"

	"github.com/oddman/mandrill/internal/cli"
	"github.com/oddman/mandrill/internal/config"
)

var banner = `
8b    d8    db    88b 88 8888b.  88""Yb 88 88     88
88b  d88   dPYb   88Yb88  8I  Yb 88__dP 88 88     88
88YbdP88  dP__Yb  88 Y88  8I  dY 88"Yb  88 88  .o 88  .o
88 YY 88 dP""""Yb 88  Y8 8888Y"  88  Yb 88 88ood8 88ood8

  Package manager & git TUI — NOT affiliated with Mailchimp's Mandrill email service.
`

func main() {
	if err := config.EnsureMandrillDir(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if !cfg.Setup.SetupComplete {
		if err := cli.FirstRunWizard(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Setup cancelled: %v\n", err)
			os.Exit(1)
		}
	}

	if err := cli.Execute(banner); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
