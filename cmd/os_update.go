package cmd

import (
	"fmt"
	"log/slog"

	helper "github.com/apexinfosysindia/cli/client"
	"github.com/spf13/cobra"
)

var osUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade", "downgrade", "up", "down"},
	Short:   "Updates the ApexOS Operating System",
	Long: `
Using this command you can upgrade or downgrade the ApexOS Operating System
to the latest version or the version specified.
`,
	Example: `
  apex os update
  apex os update --version 5
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("os update", "args", args)

		section := "os"
		command := "update"

		var options map[string]any

		version, _ := cmd.Flags().GetString("version")
		if version != "" {
			options = map[string]any{"version": version}
		}

		ProgressSpinner.Start()
		resp, err := helper.GenericJSONPostTimeout(section, command, options, helper.OsDownloadTimeout)
		ProgressSpinner.Stop()
		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else if helper.ShowJSONResponse(resp) {
			fmt.Println("\nOS update applied. Reboot the device using `apex host reboot` to finish the update.")
		} else {
			ExitWithError = true
		}
	},
}

func init() {
	osUpdateCmd.Flags().StringP("version", "", "", "Version to update to")
	osUpdateCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
	osCmd.AddCommand(osUpdateCmd)
}
