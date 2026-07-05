package cmd

import (
	"github.com/spf13/cobra"
)

var osBoardsRaspberrypiCmd = &cobra.Command{
	Use:     "raspberrypi",
	Aliases: []string{"rpi"},
	Short:   "See or change settings of the current Raspberry Pi board",
	Long: `
This command allows you to see or change settings of the Raspberry Pi board that
ApexOS is running on.`,
	Example: `
  apex os boards raspberrypi firmware`,
}

func init() {
	osBoardsCmd.AddCommand(osBoardsRaspberrypiCmd)
}
