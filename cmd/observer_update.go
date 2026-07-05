package cmd

import (
	"log/slog"

	helper "github.com/apexinfosysindia/cli/client"
	"github.com/spf13/cobra"
)

var observerUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade", "downgrade", "up", "down"},
	Short:   "Updates the internal ApexOS observer",
	Long: `
Using this command you can upgrade or downgrade the internal ApexOS 
observer, to the latest version or the version specified.
`,
	Example: `
  apex observer update --version 5
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("observer update", "args", args)

		section := "observer"
		command := "update"

		var options map[string]any

		version, _ := cmd.Flags().GetString("version")
		if version != "" {
			options = map[string]any{"version": version}
		}

		ProgressSpinner.Start()
		resp, err := helper.GenericJSONPostTimeout(section, command, options, helper.ContainerDownloadTimeout)
		ProgressSpinner.Stop()

		if err != nil {
			helper.PrintError(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}
	},
}

func init() {
	observerUpdateCmd.Flags().StringP("version", "", "", "Version to update to")
	observerUpdateCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
	observerCmd.AddCommand(observerUpdateCmd)
}
