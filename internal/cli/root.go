package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mandrill",
	Short: "A package manager and git TUI on steroids",
	Long:  "mandrill - install tools, manage repos, and navigate git like a pro.",
}

func Execute(banner string) error {
	rootCmd.SetUsageTemplate(banner + rootCmd.UsageTemplate())
	return rootCmd.Execute()
}
