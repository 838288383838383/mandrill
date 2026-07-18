package cli

import (
	"github.com/oddman/mandrill/internal/tui"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Launch the git TUI",
	Long:  "Launch the interactive git management TUI with panels for branches, files, log, and diff.",
	RunE: func(cmd *cobra.Command, args []string) error {
		theme, _ := cmd.Flags().GetString("theme")
		if theme == "" {
			theme = "catppuccin-mocha"
		}
		return tui.Run(theme)
	},
}

func init() {
	gitCmd.Flags().StringP("theme", "t", "catppuccin-mocha", "Theme to use")
	rootCmd.AddCommand(gitCmd)
}
