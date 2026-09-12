package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var coreCmd = &cobra.Command{
	Use:     "core",
	Aliases: []string{"apexos", "apexos", "apex"},
	Short:   "Provides control of the ApexOS Core",
	Long: `
This command provides a set of subcommands to control the ApexOS Core
instance running on this installation.

It provides commands to control ApexOS Core (start, stop, restart),
but also allows you to check your ApexOS Core configuration.
Furthermore, some options can be set and allows for upgrading/downgrading
ApexOS Core.
`,
	Example: `
  apex core check
  apex core restart
  apex core update
	apex core update --version 2021.11.5`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		for idx, arg := range os.Args {
			if idx != 0 && (arg == "apexos" || arg == "apex") {
				cmd.PrintErrf("The use of '%s' is deprecated, please use 'core' instead!\n", arg)
			}
		}
		rootCmd.PersistentPreRun(cmd, args)
	},
}

func init() {
	// add cmd to root command
	rootCmd.AddCommand(coreCmd)
}
