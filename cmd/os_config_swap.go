package cmd

import (
	"github.com/spf13/cobra"
)

var osConfigSwapCmd = &cobra.Command{
	Use:     "swap",
	Aliases: []string{"sw"},
	Short:   "Show or change ApexOS swap settings",
	Long: `
This command allows you to show or change current swap configuration
of ApexOS.`,
	Example: `
  apex os config swap info`,
}

func init() {
	osConfigCmd.AddCommand(osConfigSwapCmd)
}
