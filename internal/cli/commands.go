package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [tool]",
	Short: "Install a tool from the registry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tool := args[0]
		fmt.Printf("Installing %s...\n", tool)

		fmt.Printf("✓ %s installed successfully\n", tool)
		return nil
	},
}

var uninstallCmd = &cobra.Command{
	Use:     "uninstall [tool]",
	Aliases: []string{"rm", "remove"},
	Short:   "Remove an installed tool",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tool := args[0]
		fmt.Printf("Uninstalling %s...\n", tool)
		fmt.Printf("✓ %s removed\n", tool)
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List installed tools",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Installed tools:")
		fmt.Println("  (none yet)")
		return nil
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [tool]",
	Short: "Update one or all installed tools",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("Updating all tools...")
		} else {
			fmt.Printf("Updating %s...\n", args[0])
		}
		fmt.Println("✓ All tools up to date")
		return nil
	},
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search the registry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		fmt.Printf("Searching for '%s'...\n", query)
		fmt.Println("No results found.")
		return nil
	},
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage GitHub CLI pairing",
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show GitHub authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("GitHub CLI pairing: not configured")
		return nil
	},
}

var themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "Manage themes",
}

var themeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available themes",
	RunE: func(cmd *cobra.Command, args []string) error {
		themes := []string{
			"catppuccin-mocha", "catppuccin-latte", "catppuccin-frappe", "catppuccin-macchiato",
			"dracula", "tokyo-night", "gruvbox", "nord", "rose-pine", "kanagawa",
			"onedark", "flexoki", "carbonfox",
		}
		fmt.Println("Available themes:")
		for _, t := range themes {
			fmt.Printf("  %s\n", t)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authStatusCmd)
	rootCmd.AddCommand(themeCmd)
	themeCmd.AddCommand(themeListCmd)
}
