package cmd

import (
	"github.com/spf13/cobra"
)

var osConfigCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf", "cfg"},
	Short:   "Show or change ApexOS settings",
	Long: `
This command allows you to show or change settings of ApexOS.`,
	Example: `
  apex os config swap`,
}

func init() {
	osCmd.AddCommand(osConfigCmd)
}
