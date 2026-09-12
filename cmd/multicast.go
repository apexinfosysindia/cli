package cmd

import (
	"github.com/spf13/cobra"
)

var multicastCmd = &cobra.Command{
	Use:     "multicast",
	Aliases: []string{"mcast", "mc"},
	Short:   "Get information, update or configure the ApexOS Multicast",
	Long: `
The multicast command allows you to manage the internal ApexOS Multicast
backend by exposing commands to view, monitor, configure and control it.`,
	Example: `
  apex multicast info
  apex multicast update`,
}

func init() {
	rootCmd.AddCommand(multicastCmd)
}
