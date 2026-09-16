package cmd

import (
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:     "time",
	Aliases: []string{"ntp", "timedate"},
	Short:   "Get information or configure ApexOS time settings",
	Long: `
The time command allows you to view and configure the ApexOS network
time synchronization servers.`,
	Example: `
  apex time info
  apex time options --servers pool.ntp.org --fallback-servers time.google.com`,
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
