package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/838288383838383/mandrill/internal/config"
	"github.com/838288383838383/mandrill/internal/github"
)

var setupBanner = `
8b    d8    db    88b 88 8888b.  88""Yb 88 88     88
88b  d88   dPYb   88Yb88  8I  Yb 88__dP 88 88     88
88YbdP88  dP__Yb  88 Y88  8I  dY 88"Yb  88 88  .o 88  .o
88 YY 88 dP""""Yb 88  Y8 8888Y"  88  Yb 88 88ood8 88ood8

  Package manager & git TUI — NOT affiliated with Mailchimp's Mandrill email service.
`

func FirstRunWizard(cfg *config.Config) error {
	fmt.Println(setupBanner)
	fmt.Println("Welcome to mandrill! First-run setup detected.")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	ghClient := &github.GhClient{}
	ghAvail := ghClient.IsAvailable()

	if ghAvail {
		fmt.Print("? Pair with GitHub CLI (gh)? [y/n/fuckoff] ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		switch answer {
		case "y", "yes":
			fmt.Println("\n  Checking gh auth status...")
			detected, err := github.DetectGhCli()
			if err != nil {
				fmt.Printf("  ✗ Error detecting gh: %v\n", err)
			} else if detected.Authenticated {
				fmt.Printf("  ✓ Logged in as %s (github.com)\n", detected.Username)
				cfg.GitHub.Enabled = true
				cfg.GitHub.PairWithGh = true
				cfg.GitHub.Username = detected.Username
			} else {
				fmt.Println("  ✗ Not logged in to gh. Skipping pairing.")
				cfg.GitHub.Enabled = false
				cfg.GitHub.PairWithGh = false
			}

		case "n", "no":
			fmt.Println("\n  Skipping GitHub pairing.")
			cfg.GitHub.Enabled = false
			cfg.GitHub.PairWithGh = false

		case "fuckoff":
			fmt.Println("\n  Understood. Skipping all prompts.")
			cfg.Setup.SetupComplete = true
			cfg.GitHub.Enabled = false
			cfg.GitHub.PairWithGh = false
			return config.Save(cfg)

		default:
			fmt.Println("\n  Invalid input. Skipping GitHub pairing.")
			cfg.GitHub.Enabled = false
			cfg.GitHub.PairWithGh = false
		}
	} else {
		fmt.Println("  GitHub CLI (gh) not found. Skipping GitHub integration.")
		cfg.GitHub.Enabled = false
		cfg.GitHub.PairWithGh = false
	}

	fmt.Printf("\n? Install prefix [%s]: ", cfg.Install.Prefix)
	prefix, _ := reader.ReadString('\n')
	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		cfg.Install.Prefix = prefix
	}

	fmt.Printf("? Default theme [%s]: ", cfg.Theme.Name)
	theme, _ := reader.ReadString('\n')
	theme = strings.TrimSpace(theme)
	if theme != "" {
		cfg.Theme.Name = theme
	}

	cfg.Setup.SetupComplete = true
	cfg.Setup.FirstRunAt = time.Now().Format(time.RFC3339)

	fmt.Println("\n  Setup complete! Run `mandrill install <tool>` to get started.")
	return config.Save(cfg)
}
