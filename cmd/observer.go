package cmd

import (
	"github.com/spf13/cobra"
)

var observerCmd = &cobra.Command{
	Use:   "observer",
	Short: "Get information, update or configure the ApexOS observer",
	Long: `
The observer command allows you to manage the internal ApexOS observer by
exposing commands to view, monitor, configure and control it.`,
	Example: `
  apex observer info
  apex observer update`,
}

func init() {
	rootCmd.AddCommand(observerCmd)
}
