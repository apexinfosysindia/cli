package cmd

import (
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Get information, update or configure the ApexOS cli backend",
	Long: `
The cli command allows you to manage the internal ApexOS CLI backend by
exposing commands to view, monitor, configure and control it.`,
	Example: `
  apex cli info
  apex cli update`,
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
