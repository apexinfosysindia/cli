package cmd

import (
	"log/slog"

	helper "github.com/apexinfosysindia/cli/client"
	"github.com/spf13/cobra"
)

var timeInfoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"in", "inf"},
	Short:   "Shows ApexOS time configuration",
	Long: `
Shows the configured ApexOS NTP servers.`,
	Example: `
  apex time info`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("time info", "args", args)

		section := "time"
		command := "info"

		resp, err := helper.GenericJSONGet(section, command)
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	timeCmd.AddCommand(timeInfoCmd)
}
