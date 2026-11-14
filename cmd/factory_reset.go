package cmd

import (
	"github.com/spf13/cobra"
)

var factoryResetCmd = &cobra.Command{
	Use:   "factory-reset",
	Short: "Wipe all ApexOS user data (alias for 'apex os datadisk wipe')",
	Long:  osDataDiskWipeCmd.Long,
	Example: `
  apex factory-reset
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run:               osDataDiskWipeCmd.Run,
}

func init() {
	rootCmd.AddCommand(factoryResetCmd)
}
