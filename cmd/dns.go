package cmd

import (
	"github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Get information, update or configure the ApexOS DNS server",
	Long: `
The dns command allows you to manage the internal ApexOS DNS server by
exposing commands to view, monitor, configure and control it.`,
	Example: `
  apex dns logs
  apex dns info
  apex dns update`,
}

func init() {
	rootCmd.AddCommand(dnsCmd)
}
