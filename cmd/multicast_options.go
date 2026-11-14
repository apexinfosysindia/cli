package cmd

import (
	"log/slog"

	helper "github.com/apexinfosysindia/cli/client"
	"github.com/spf13/cobra"
)

var multicastOptionsCmd = &cobra.Command{
	Use:     "options",
	Aliases: []string{"option", "opt", "opts", "op"},
	Short:   "Allow to set options for the internal ApexOS Multicast server",
	Long: `
This command allows you to set configuration options for the internally
running ApexOS Multicast server.
`,
	Example: `
  apex multicast options --enabled=false
  apex multicast options --enabled
`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Args:              cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("multicast options", "args", args)

		section := "multicast"
		command := "options"

		options := make(map[string]any)

		enabled, err := cmd.Flags().GetBool("enabled")
		if err == nil && cmd.Flags().Changed("enabled") {
			options["enabled"] = enabled
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
	multicastOptionsCmd.Flags().BoolP("enabled", "", true, "Enable/Disable the Multicast plug-in")

	multicastOptionsCmd.Flags().Lookup("enabled").NoOptDefVal = "true"

	multicastOptionsCmd.RegisterFlagCompletionFunc("enabled", boolCompletions)

	multicastCmd.AddCommand(multicastOptionsCmd)
}
